package infrastructure

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	Client() *redis.Client
	Incr(ctx context.Context, key string) *redis.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
}

type RealRedisClient struct {
	client *redis.Client
}

func NewRedisClient(redisAddress string) *RealRedisClient {
	return &RealRedisClient{
		client: redis.NewClient(&redis.Options{
			Addr: redisAddress,
		}),
	}
}

func (r *RealRedisClient) Client() *redis.Client {
	return r.client
}

func (r *RealRedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	return r.client.Incr(ctx, key)
}

func (r *RealRedisClient) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.Expire(ctx, key, expiration)
}
