package redis

import (
	"time"

	"github.com/pkg/errors"

	"github.com/go-redis/redis"
)

var DB *redis.Client

func init() {
	DB = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	DB.WrapProcess(func(old func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			// TODO(shijiada): add tracing
			err := old(cmd)
			return errors.Wrapf(err, "failed to processing <%s>", cmd)
		}
	})
}
