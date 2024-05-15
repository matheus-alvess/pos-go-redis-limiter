package port

import (
	"context"
	"time"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository_mock.go -package=port
type RateLimiterRepository interface {
	Allow(ctx context.Context, key string, limit int, duration time.Duration) (bool, error)
}
