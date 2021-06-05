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
	err := r.rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *RedisCache) Get(key string, defaultValue string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return defaultValue, nil
	} else if err != nil {
		return "", err
	}

	return val, nil
}
