package infrastructure

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(config *Config) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
	})

	return &RedisClient{
		client: client,
	}
}

func (r *RedisClient) Allow(ctx context.Context, key string, limit int) (*redis_rate.Result, error) {
	panic("")
}
