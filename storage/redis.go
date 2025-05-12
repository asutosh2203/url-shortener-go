package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// SET function wrapper
func Set(key string, value string, expr time.Duration) error {
	return RedisClient.Set(Ctx, key, value, expr).Err()
}

// GET function wrapper
func Get(key string) (string, error) {
	return RedisClient.Get(Ctx, key).Result()
}

// GET the TTL of a key
func GetTTL(key string) (time.Duration, error) {
	return RedisClient.TTL(Ctx, key).Result()
}
