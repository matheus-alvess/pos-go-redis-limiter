package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func setEnv(key, value string) {
	os.Setenv(key, value)
}

func unsetEnv(key string) {
	os.Unsetenv(key)
}

func TestLoad(t *testing.T) {
	t.Run("should load config successfully", func(t *testing.T) {
		setEnv("REDIS_ADDR", "localhost:6379")
		setEnv("IP_RATE_LIMIT_PER_SECOND", "10")
		setEnv("TOKEN_RATE_LIMIT_PER_SECOND", "20")
		setEnv("GENERAL_TIME_BAN", "1h")

		config, err := Load()
		require.NoError(t, err)

		assert.Equal(t, "localhost:6379", config.RedisAddress)
		assert.Equal(t, 10, config.IpRateLimitPerSecond)
		assert.Equal(t, 20, config.TokenRateLimitPerSecond)
		assert.Equal(t, time.Hour, config.GeneralTimeBan)

		unsetEnv("REDIS_ADDR")
		unsetEnv("IP_RATE_LIMIT_PER_SECOND")
		unsetEnv("TOKEN_RATE_LIMIT_PER_SECOND")
		unsetEnv("GENERAL_TIME_BAN")
	})

	t.Run("should return error if REDIS_ADDR is not set", func(t *testing.T) {
		_, err := Load()
		assert.EqualError(t, err, "REDIS_ADDR config notfound")
	})

	t.Run("should return error if IP_RATE_LIMIT_PER_SECOND is not set", func(t *testing.T) {
		setEnv("REDIS_ADDR", "localhost:6379")
		_, err := Load()
		assert.EqualError(t, err, "IP_RATE_LIMIT_PER_SECOND config notfound")
		unsetEnv("REDIS_ADDR")
	})

	t.Run("should return error if TOKEN_RATE_LIMIT_PER_SECOND is not set", func(t *testing.T) {
		setEnv("REDIS_ADDR", "localhost:6379")
		setEnv("IP_RATE_LIMIT_PER_SECOND", "10")
		_, err := Load()
		assert.EqualError(t, err, "TOKEN_RATE_LIMIT_PER_SECOND config notfound")
		unsetEnv("REDIS_ADDR")
		unsetEnv("IP_RATE_LIMIT_PER_SECOND")
	})

	t.Run("should return error if GENERAL_TIME_BAN is not set", func(t *testing.T) {
		setEnv("REDIS_ADDR", "localhost:6379")
		setEnv("IP_RATE_LIMIT_PER_SECOND", "10")
		setEnv("TOKEN_RATE_LIMIT_PER_SECOND", "20")
		_, err := Load()
		assert.EqualError(t, err, "GENERAL_TIME_BAN config notfound")
		unsetEnv("REDIS_ADDR")
		unsetEnv("IP_RATE_LIMIT_PER_SECOND")
		unsetEnv("TOKEN_RATE_LIMIT_PER_SECOND")
	})

	t.Run("should return error if IP_RATE_LIMIT_PER_SECOND is not an integer", func(t *testing.T) {
		setEnv("REDIS_ADDR", "localhost:6379")
		setEnv("IP_RATE_LIMIT_PER_SECOND", "not-an-integer")
		_, err := Load()
		assert.Error(t, err)
		unsetEnv("REDIS_ADDR")
		unsetEnv("IP_RATE_LIMIT_PER_SECOND")
	})

	t.Run("should return error if TOKEN_RATE_LIMIT_PER_SECOND is not an integer", func(t *testing.T) {
		setEnv("REDIS_ADDR", "localhost:6379")
		setEnv("IP_RATE_LIMIT_PER_SECOND", "10")
		setEnv("TOKEN_RATE_LIMIT_PER_SECOND", "not-an-integer")
		_, err := Load()
		assert.Error(t, err)
		unsetEnv("REDIS_ADDR")
		unsetEnv("IP_RATE_LIMIT_PER_SECOND")
		unsetEnv("TOKEN_RATE_LIMIT_PER_SECOND")
	})

	t.Run("should return error if GENERAL_TIME_BAN is not a valid duration", func(t *testing.T) {
		setEnv("REDIS_ADDR", "localhost:6379")
		setEnv("IP_RATE_LIMIT_PER_SECOND", "10")
		setEnv("TOKEN_RATE_LIMIT_PER_SECOND", "20")
		setEnv("GENERAL_TIME_BAN", "invalid-duration")
		_, err := Load()
		assert.Error(t, err)
		unsetEnv("REDIS_ADDR")
		unsetEnv("IP_RATE_LIMIT_PER_SECOND")
		unsetEnv("TOKEN_RATE_LIMIT_PER_SECOND")
		unsetEnv("GENERAL_TIME_BAN")
	})
}
