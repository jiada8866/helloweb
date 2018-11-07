package redis

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	// TODO 暂时通过全局变量来初始化
	Pool *redis.Pool
)

func init() {
	// redis 地址可以从环境变量中获取
	redisHost := os.Getenv("POPUPS_REDIS_HOST")
	if redisHost == "" {
		// 默认连接 localhost:6379
		redisHost = ":6379"
	}
	Pool = newPool(redisHost)
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
