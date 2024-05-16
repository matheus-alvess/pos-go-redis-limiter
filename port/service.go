package port

import (
	"time"
)

//go:generate mockgen -source=./service.go -destination=./mock/service_mock.go -package=mocks
type RateLimiterService interface {
	Allow(key string, limit int, duration time.Duration) (bool, error)
}
