package port

import (
	"context"
	"github.com/go-redis/redis_rate/v9"
)

type RateLimiterRepository interface {
	Allow(ctx context.Context, key string, secondsInterval int) (*redis_rate.Result, error)
}

type RateLimiterService interface {
	Allow(ctx context.Context, ip string, secondsInterval int) error
}
