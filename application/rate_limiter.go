package application

import (
	"context"
	"pos-go-redis-limiter/port"
	"time"
)

type rateLimiterServiceHandler struct {
	rateLimiterRepository port.RateLimiterRepository
}

func NewRateLimiterService(rateLimiterRepository port.RateLimiterRepository) port.RateLimiterService {
	return &rateLimiterServiceHandler{
		rateLimiterRepository: rateLimiterRepository,
	}
}

func (r *rateLimiterServiceHandler) Allow(ctx context.Context, key string, limit int, duration time.Duration) (bool, error) {
	return r.rateLimiterRepository.Allow(ctx, key, limit, duration)
}
