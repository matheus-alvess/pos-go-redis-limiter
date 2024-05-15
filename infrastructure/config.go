package infrastructure

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	RedisAddress       string
	RateLimitPerSecond int
	BandDuration       string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	redisAddress, ok := os.LookupEnv("REDIS_ADDR")
	if !ok {
		return nil, errors.New("REDIS_ADDR config notfound")
	}

	rateLimitPerSecond, ok := os.LookupEnv("RATE_LIMIT_PER_SECOND")
	if !ok {
		return nil, errors.New("RATE_LIMIT_PER_SECOND config notfound")
	}

	rateLimitPerSecondValue, err := strconv.Atoi(rateLimitPerSecond)
	if err != nil {
		return nil, err
	}

	banDuration, ok := os.LookupEnv("BAN_DURATION")
	if !ok {
		return nil, errors.New("BAN_DURATION config notfound")
	}

	return &Config{
		RedisAddress:       redisAddress,
		RateLimitPerSecond: rateLimitPerSecondValue,
		BandDuration:       banDuration,
	}, nil
}
