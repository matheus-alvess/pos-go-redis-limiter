package port

import (
	"context"
	"time"
)

//go:generate mockgen -source=./service.go -destination=./mock/service_mock.go -package=port
type RateLimiterService interface {
	Allow(ctx context.Context, key string, limit int, duration time.Duration) (bool, error)
}
