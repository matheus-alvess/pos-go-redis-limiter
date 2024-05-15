package infrastructure

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	RedisAddress            string
	IpRateLimitPerSecond    int
	TokenRateLimitPerSecond int
	GeneralTimeBan          time.Duration
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

	ipRateLimitPerSecond, ok := os.LookupEnv("IP_RATE_LIMIT_PER_SECOND")
	if !ok {
		return nil, errors.New("IP_RATE_LIMIT_PER_SECOND config notfound")
	}

	ipRateLimitPerSecondValue, err := strconv.Atoi(ipRateLimitPerSecond)
	if err != nil {
		return nil, err
	}

	tokenRateLimitPerSecond, ok := os.LookupEnv("TOKEN_RATE_LIMIT_PER_SECOND")
	if !ok {
		return nil, errors.New("TOKEN_RATE_LIMIT_PER_SECOND config notfound")
	}

	tokenRateLimitPerSecondValue, err := strconv.Atoi(tokenRateLimitPerSecond)
	if err != nil {
		return nil, err
	}

	generalTimeBan, ok := os.LookupEnv("GENERAL_TIME_BAN")
	if !ok {
		return nil, errors.New("GENERAL_TIME_BAN config notfound")
	}

	generalTimeBanValue, err := time.ParseDuration(generalTimeBan)
	if err != nil {
		return nil, err
	}

	return &Config{
		RedisAddress:            redisAddress,
		IpRateLimitPerSecond:    ipRateLimitPerSecondValue,
		TokenRateLimitPerSecond: tokenRateLimitPerSecondValue,
		GeneralTimeBan:          generalTimeBanValue,
	}, nil
}
