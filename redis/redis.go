package redis

import (
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

// Connect подключение к key-value хранилищу redis
func Connect(storageUser, storagePass, storageHost, database string) (*redis.Client, error) {
	databaseID, err := strconv.Atoi(database)
	if err != nil {
		return nil, errors.Wrap(err, "database must be number")
	}

	redisCli := redis.NewClient(&redis.Options{
		Addr:     storageHost,
		Password: storagePass,
		DB:       databaseID,
	})
	_, err = redisCli.Ping().Result()
	if err != nil {
		return nil, errors.Wrap(err, "can not ping redis server")
	}

	return redisCli, nil
}
