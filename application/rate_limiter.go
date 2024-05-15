package application

import (
	"context"
	"log"
	"pos-go-redis-limiter/port"
)

type rateLimiterServiceHandler struct {
	rateLimiterRepository port.RateLimiterRepository
}

func NewRateLimiterService(rateLimiterRepository port.RateLimiterRepository) port.RateLimiterService {
	return &rateLimiterServiceHandler{
		rateLimiterRepository: rateLimiterRepository,
	}
}

func (r rateLimiterServiceHandler) Allow(ctx context.Context, ip string, secondsInterval int) error {
	res, err := r.rateLimiterRepository.Allow(ctx, ip, secondsInterval)
	if err != nil {
		log.Printf("Rate limit Error: %s", err.Error())
		return err
	}

	log.Printf("Request Received for IP: %s, ALLOWED: %d, REMAINING: %d", ip, res.Allowed, res.Remaining)

	if res.Allowed == 0 {
		return ErrLimitExceeded
	}

	return nil
}
