package redis

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type Locker struct {
	client *redis.Client

	key    string
	value  string
	expiry time.Duration
}

func NewLocker(client *redis.Client, key string, expiry time.Duration) *Locker {
	return &Locker{
		client: client,

		key:    key + ":lock",
		expiry: expiry,
	}
}

func (l *Locker) Lock() (bool, error) {
	value, err := genValue()
	if err != nil {
		return false, errors.Wrap(err, "failed to genValue")
	}

	l.value = value

	reply, err := l.client.SetNX(l.key, value, l.expiry).Result()
	if err != nil && err != redis.Nil {
		return false, errors.Wrap(err, "failed to lock")
	}
	return reply, nil
}

var deleteScript = redis.NewScript(`
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	else
		return 0
	end
`)

func (l *Locker) Unlock() (bool, error) {
	status, err := deleteScript.Run(l.client, []string{l.key}, l.value).Result()
	if err != nil {
		return false, errors.Wrap(err, "failed to unlock")
	}
	return status != 0, nil
}

func genValue() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
