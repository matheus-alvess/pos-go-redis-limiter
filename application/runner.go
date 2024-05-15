package application

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pos-go-redis-limiter/infrastructure"
)

func StartupApp(config *infrastructure.Config) *gin.Engine {
	redisRepo := infrastructure.NewRateLimiterRepository(config)
	rateLimiterService := NewRateLimiterService(redisRepo)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(RateLimitMiddleware(rateLimiterService, config.RateLimitPerSecond))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello FullCycle",
		})
	})

	return r
}
