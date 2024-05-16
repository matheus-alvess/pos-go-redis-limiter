package infrastructure

import (
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
	"pos-go-redis-limiter/port"
	"time"
)

type rateLimiterRepositoryHandler struct {
	client       redis.Cmdable
	redisLimiter *redis_rate.Limiter
}

func NewRateLimiterRepository(client redis.Cmdable) port.RateLimiterRepository {
	limiter := redis_rate.NewLimiter(client)

	return &rateLimiterRepositoryHandler{
		client:       client,
		redisLimiter: limiter,
	}
}

func (r *rateLimiterRepositoryHandler) Allow(key string, limit int, duration time.Duration) (bool, error) {
	val, err := r.client.Incr(key).Result()
	if err != nil {
		return false, err
	}
	
	if val == 1 {
		r.client.Expire(key, duration)
	}

	return val <= int64(limit), nil
}
