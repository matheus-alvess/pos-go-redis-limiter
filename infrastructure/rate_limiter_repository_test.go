package infrastructure

import (
	"errors"
	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"log"
	"pos-go-redis-limiter/port"
	"testing"
	"time"
)

type RepositoryTestSuite struct {
	suite.Suite
	repository      port.RateLimiterRepository
	redisClientMock *redis.Client
}

func (suite *RepositoryTestSuite) SetupTest() {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	suite.redisClientMock = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestAllow() {
	suite.Run("should execute allow with not expired key", func() {
		key := "123"
		limit := 2
		duration := time.Second * 1

		currentRedisMock := redismock.NewNiceMock(suite.redisClientMock)

		mockRedisCMD := redis.NewIntResult(int64(2), nil)

		currentRedisMock.On("Incr", mock.Anything).Return(mockRedisCMD)

		repository := NewRateLimiterRepository(currentRedisMock)

		ok, err := repository.Allow(key, limit, duration)
		assert.Nil(suite.T(), err)
		assert.True(suite.T(), ok)
	})

	suite.Run("should execute allow with expired key", func() {
		key := "123"
		limit := 1
		duration := time.Second * 1

		currentRedisMock := redismock.NewNiceMock(suite.redisClientMock)

		mockRedisCMD := redis.NewIntResult(int64(1), nil)
		mockRedisBoolResult := redis.NewBoolResult(true, nil)

		currentRedisMock.On("Incr", mock.Anything).Return(mockRedisCMD)
		currentRedisMock.On("Expire", mock.Anything, mock.Anything).Return(mockRedisBoolResult)

		repository := NewRateLimiterRepository(currentRedisMock)

		ok, err := repository.Allow(key, limit, duration)
		assert.Nil(suite.T(), err)
		assert.True(suite.T(), ok)
	})

	suite.Run("should execute allow with failed", func() {
		key := "123"
		limit := 2
		duration := time.Second * 1

		currentRedisMock := redismock.NewNiceMock(suite.redisClientMock)

		expectedErr := errors.New("mock error")
		mockRedisCMD := redis.NewIntResult(int64(0), expectedErr)

		currentRedisMock.On("Incr", mock.Anything).Return(mockRedisCMD)

		repository := NewRateLimiterRepository(currentRedisMock)

		_, err := repository.Allow(key, limit, duration)
		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "mock error", err.Error())
	})
}
