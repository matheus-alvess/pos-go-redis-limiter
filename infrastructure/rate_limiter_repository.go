package infrastructure

import (
	"context"
	"github.com/go-redis/redis_rate/v9"
)

type RateLimiterRepositoryImpl struct {
	redisClient  *RedisClient
	redisLimiter *redis_rate.Limiter
}

func NewRateLimiterRepository(config *Config) *RateLimiterRepositoryImpl {
	redisClient := NewRedisClient(config)
	limiter := redis_rate.NewLimiter(redisClient.client)

	return &RateLimiterRepositoryImpl{
		redisClient:  redisClient,
		redisLimiter: limiter,
	}
}

func (r *RateLimiterRepositoryImpl) Allow(ctx context.Context, key string, secondsInterval int) (*redis_rate.Result, error) {
	return r.redisLimiter.Allow(ctx, key, redis_rate.PerMinute(secondsInterval))
}
