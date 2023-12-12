package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisCli *redis.Client
var ctx = context.Background()

func Start() error {
	if redisCli == nil {
		redisCli = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		fmt.Println("Redis Client Listening")

		redisCli.Set(ctx, "test", "test", 0).Result()
	}
	return nil
}

func Get(key string) (string, error) {
	value, err := redisCli.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("keys not exist")
	} else if err != nil {
		return "", err
	}
	return value, nil
}

func Delete(key string) error {
	return redisCli.Del(ctx, key).Err()
}

func Set(key, value string) error {
	return redisCli.Set(ctx, key, value, 0).Err()
}
