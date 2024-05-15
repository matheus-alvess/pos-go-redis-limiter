package infrastructure

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"pos-go-redis-limiter/infrastructure/mocks"
	"pos-go-redis-limiter/port"
	"testing"
	"time"
)

type RepositoryTestSuite struct {
	suite.Suite
	context         context.Context
	repository      port.RateLimiterRepository
	redisClientMock *mocks.MockRedisClient
}

func (suite *RepositoryTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.redisClientMock = mocks.NewMockRedisClient(ctrl)
	suite.context = context.Background()
	suite.repository = NewRateLimiterRepository(suite.redisClientMock)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestNewRepository() {
	suite.Run("success instantiate Repository", func() {

		repository := NewRateLimiterRepository(suite.redisClientMock)
		assert.NotNil(suite.T(), repository)
	})
}

func (suite *RepositoryTestSuite) TestAllow() {
	key := ""
	limit := 2
	duration := time.Second * 1

	suite.Run("should execute allow with success", func() {
		suite.redisClientMock.EXPECT().Client().Return(&redis.Client{}).AnyTimes()

		ok, err := suite.repository.Allow(suite.context, key, limit, duration)

		assert.Nil(suite.T(), err)
		assert.True(suite.T(), ok)
	})
}
