package redis

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

// Connect подключение к key-value хранилищу redis
func Connect(storageUser, storagePass, storageHost string) (*redis.Client, error) {
	var err error
	redisCli := redis.NewClient(&redis.Options{
		Addr:     storageHost,
		Password: storagePass,
		DB:       0,
	})
	_, err = redisCli.Ping().Result()
	if err != nil {
		return nil, errors.Wrap(err, "can not ping redis server")
	}

	return redisCli, nil
}
