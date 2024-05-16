package application

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"pos-go-redis-limiter/infrastructure"
)

func StartupApp(config *infrastructure.Config) *gin.Engine {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
	})
	redisRepo := infrastructure.NewRateLimiterRepository(client)
	rateLimiterService := NewRateLimiterService(redisRepo)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(RateLimitMiddleware(rateLimiterService, config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello FullCycle",
		})
	})

	return r
}
