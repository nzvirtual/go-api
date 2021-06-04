package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	rdb *redis.Client
}

var ctx = context.Background()

func ConnectRedis(options *Options) (*RedisCache, error) {
	rc := &RedisCache{}
	rc.rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", options.Hostname, options.Port),
		Password: options.Password,
		DB:       0,
	})

	return rc, nil
}

func (r *RedisCache) Set(key string, value string) (bool, error) {
	return true, nil
}

func (r *RedisCache) Get(key string, defaultValue string) (string, error) {
	return "", nil
}
