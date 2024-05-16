package port

import (
	"time"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository_mock.go -package=mocks
type RateLimiterRepository interface {
	Allow(key string, limit int, duration time.Duration) (bool, error)
}
