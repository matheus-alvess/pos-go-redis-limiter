package application

import (
	"errors"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"pos-go-redis-limiter/infrastructure"
	mocks "pos-go-redis-limiter/port/mock"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRateLimitMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	gin.SetMode(gin.TestMode)

	mockService := mocks.NewMockRateLimiterService(ctrl)

	config := &infrastructure.Config{
		IpRateLimitPerSecond:    2,
		TokenRateLimitPerSecond: 5,
		GeneralTimeBan:          time.Second * 10,
	}

	t.Run("should allow request without API token", func(t *testing.T) {
		mockService.EXPECT().Allow(gomock.Any(), config.IpRateLimitPerSecond, config.GeneralTimeBan).Return(true, nil)

		router := gin.New()
		router.Use(RateLimitMiddleware(mockService, config))
		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello FullCycle")
		})

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = "127.0.0.1:12345"
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Hello FullCycle", w.Body.String())
	})

	t.Run("should allow request with API token", func(t *testing.T) {
		mockService.EXPECT().Allow(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)

		router := gin.New()
		router.Use(RateLimitMiddleware(mockService, config))
		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello FullCycle")
		})

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("API_KEY", "test-token")
		req.RemoteAddr = "127.0.0.1:12345"
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Hello FullCycle", w.Body.String())
	})

	t.Run("should block request if rate limit is exceeded", func(t *testing.T) {
		mockService.EXPECT().Allow(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)

		router := gin.New()
		router.Use(RateLimitMiddleware(mockService, config))
		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusTooManyRequests, "Rate limit exceeded")
		})

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = "127.0.0.1:12345"
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusTooManyRequests, w.Code)
		assert.Equal(t, "you have reached the maximum number of requests or actions allowed within a certain time frame", w.Body.String())
	})

	t.Run("should return internal server error on rate limiter error", func(t *testing.T) {
		mockService.EXPECT().Allow(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, errors.New("rate limiter error"))

		router := gin.New()
		router.Use(RateLimitMiddleware(mockService, config))
		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusInternalServerError, "Internal server error")
		})

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = "127.0.0.1:12345"
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "rate limiter error", w.Body.String())
	})
}
