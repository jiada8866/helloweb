package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

func main() {
	logger := logrus.New()
	tf := new(logrus.TextFormatter)
	tf.TimestampFormat = time.RFC3339Nano
	logger.Formatter = tf
	logrus.SetFormatter(tf)

	logfile, err := os.Create("/tmp/log/rate.log")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer logfile.Close()
	logger.SetOutput(logfile)

	limiter := rate.NewLimiter(rate.Limit(50), 25)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(ctx context.Context, index int) {
			defer wg.Done()

			i := inquirer{
				request: gorequest.New(),
				limiter: limiter,
				logger:  logger,
				onRate:  true,
			}

			for {
				select {
				case <-ctx.Done():
					logrus.WithField("index", index).WithError(ctx.Err()).Info("time is up")
					return
				default:
					_, err := i.query(ctx, "http://127.0.0.1:25378/one/", index)
					if err != nil {
						logrus.WithError(err).WithField("index", index).Error("failed to query")
					}
				}
			}
		}(ctx, j)
	}

	wg.Wait()
}

type inquirer struct {
	request *gorequest.SuperAgent
	limiter *rate.Limiter
	logger  *logrus.Logger
	onRate  bool
}

// 对 client 端在访问 server 时做限速，有以下几种情况：
// 	1.client 的访问速度比 server 的极限 qps 还要高，不能让 client 开全速
// 	2.client 的访问速度比 server 的极限 qps 要小，限速以防 client 在某些时刻突然抽风变快（网络突然变畅通导致速度变快）
// Wait 方法接收带有超时时间的 context 的话，如果在超时时间内等不到 token 会直接报错返回，
// 所以调用 Wait 的调用方需要自己判断是否需要重试！
// 不用带超时的 context 的话，Wait 方法会一直等到有 token 再往下执行，即 block 相应的 goroutine
func (i *inquirer) query(ctx context.Context, url string, index int) ([]byte, error) {
	u1 := uuid.NewV4()
	contextLogger := i.logger.WithFields(logrus.Fields{"uuid": u1, "index": fmt.Sprintf("index%d", index)})

	contextLogger.Info("before wait")
	if i.onRate {
		// 不用方法调用者传过来的 ctx，而用 context.TODO() 来表示可以无限期等到有可用的 token 为止
		if err := i.limiter.Wait(context.TODO()); err != nil {
			return nil, err
		}
		contextLogger.Info("after wait")
	}

	resp, body, errs := i.request.Get(url + u1.String()).EndBytes()
	if errs != nil {
		return nil, fmt.Errorf("request: %s", errs)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code:%d", resp.StatusCode)
	}

	return body, nil
}
