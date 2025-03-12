package utilities

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type RedisService struct {
	client *redis.Client
}

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis-dev.redis-dev:6379", // "localhost:6379", //
		Password: "",
		DB:       0,
	})
}

func RedisSetCache(key string, value interface{}, expiresAt time.Duration) error {
	client := getClient()

	// serialize the post
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	return client.Set(ctx, key, json, expiresAt).Err()
}

func RedisGetCache(key string) (string, error) {
	client := getClient()

	value, err := client.Get(ctx, key).Result()

	if err != nil {
		return "nil", errors.New("RedisError: Cannot find any matching key")
	}

	return value, nil
}

func RedisDeleteKey(key string) error {
	client := getClient()

	_, getCacheErr := RedisGetCache(key)

	if getCacheErr != nil {
		return getCacheErr
	}

	return client.Del(ctx, key).Err()
}
