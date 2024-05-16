package application

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"pos-go-redis-limiter/port"
	mocks "pos-go-redis-limiter/port/mock"
	"testing"
	"time"
)

type ServiceTestSuite struct {
	suite.Suite
	mockRepository *mocks.MockRateLimiterRepository
	service        port.RateLimiterService
}

func (suite *ServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.mockRepository = mocks.NewMockRateLimiterRepository(ctrl)
	suite.service = NewRateLimiterService(suite.mockRepository)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) TestAllow() {
	suite.Run("should execute allow success", func() {
		key := "123"
		limit := 2
		duration := time.Second * 1

		suite.mockRepository.EXPECT().Allow(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)

		ok, err := suite.service.Allow(key, limit, duration)
		assert.Nil(suite.T(), err)
		assert.True(suite.T(), ok)
	})

	suite.Run("should execute allow failed", func() {
		key := "123"
		limit := 2
		duration := time.Second * 1

		suite.mockRepository.EXPECT().Allow(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, errors.New("mock error"))

		ok, err := suite.service.Allow(key, limit, duration)
		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "mock error", err.Error())
		assert.False(suite.T(), ok)
	})
}
