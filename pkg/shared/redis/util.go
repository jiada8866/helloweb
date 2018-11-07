package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func Ping() error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return errors.Wrap(err, "cannot 'PING' db")
	}
	return nil
}

func Get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, errors.Wrapf(err, "failed to get key %s", key)
	}
	return data, nil
}

func Set(key string, value []byte) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return errors.Wrapf(err, "failed to set key %s to %s", key, v)
	}
	return nil
}

func SetEX(key string, seconds int64, value []byte) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, seconds, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return errors.Wrapf(err, "failed to set key %s to %s with expire %d seconds", key, v, seconds)
	}
	return nil
}

func Exists(key string) (bool, error) {
	conn := Pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, errors.Wrapf(err, "failed to check if key %s exists", key)
	}
	return ok, nil
}

func Delete(key string) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	if err != nil {
		return errors.Wrapf(err, "failed to del key %s", key)
	}
	return nil
}

func GetKeys(pattern string) ([]string, error) {
	conn := Pool.Get()
	defer conn.Close()

	iter := 0
	var keys []string
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, errors.Wrapf(err, "failed to retrieve '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

func Incr(counterKey string) (int, error) {
	conn := Pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("INCR", counterKey))
	if err != nil {
		return i, errors.Wrapf(err, "failed to incr key %s", counterKey)
	}
	return i, nil
}
