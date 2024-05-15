package infrastructure

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	"pos-go-redis-limiter/port"
	"time"
)

type rateLimiterRepositoryHandler struct {
	client       *redis.Client
	redisLimiter *redis_rate.Limiter
}

func NewRateLimiterRepository(instanceWrapperRedis RedisClient) port.RateLimiterRepository {
	redisClient := instanceWrapperRedis.Client()
	limiter := redis_rate.NewLimiter(redisClient)

	return &rateLimiterRepositoryHandler{
		client:       redisClient,
		redisLimiter: limiter,
	}
}

func (r *rateLimiterRepositoryHandler) Allow(ctx context.Context, key string, limit int, duration time.Duration) (bool, error) {
	val, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if val == 1 {
		r.client.Expire(ctx, key, duration)
	}

	return val <= int64(limit), nil
}
