// Package echologrus provides a middleware for echo that logs request details
// via the logrus logging library
package middleware

import (
	"os"
	"time"

	"github.com/jiadas/helloweb/pkg/shared/log"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

var hostname string

func init() {
	hostname, _ = os.Hostname()
}

func AccessLog() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			start := time.Now()

			if err := next(c); err != nil {
				c.Error(err)
			}

			latency := time.Since(start)

			bytesIn := req.Header.Get(echo.HeaderContentLength)
			if bytesIn == "" {
				bytesIn = "0"
			}

			log.G(req.Context()).WithFields(logrus.Fields{
				"type":          "access",
				"hostname":      hostname,
				"remote_ip":     c.RealIP(),
				"host":          req.Host,
				"uri":           req.RequestURI,
				"method":        req.Method,
				"status":        res.Status,
				"route":         c.Path(),
				"latency":       int64(latency),
				"latency_human": latency,
				"bytes_in":      bytesIn,
				"bytes_out":     res.Size,
			}).Info("request access log info")

			// c.Error(err) 会在好多 echo 自带的中间件中使用
			// 但是这些自带的中间件在当 next(c) 返回的 err != nil 时调用 c.Error(err) 同时 return err
			// 这样会导致同一个 err 在 ErrorHandler 中被重复处理多次
			//
			// 可以把 AccessLog 这个中间件放在一些列 Use 的最后，因为这里的 return nil 相当于让后续中间件的 next(c) 的返回值都是 nil
			// 就可以不重复触发 c.Error(err)
			// 但是依靠固定中间件的顺序来解决重复处理错误还是显得不够彻底
			return nil
		}
	}
}
