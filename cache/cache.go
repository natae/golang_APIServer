package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisClient *redis.Client = nil
var ctx = context.Background()

func getRdb() *redis.Client {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
			Password: "",
			DB: 0,
		})
	}
	return redisClient
}

func Set(key string, val string, expire time.Duration) (bool, error) {
	rdb := getRdb()
	err := rdb.Set(ctx, key, val, expire).Err()
	if err != nil {
		panic(err)
		return false, err
	}
	return true, err
}

func Get(key string) (string, error) {
	rdb := getRdb()
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
		return "", err
	}
	return val, err
}

