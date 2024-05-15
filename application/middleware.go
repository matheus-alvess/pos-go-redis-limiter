package application

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pos-go-redis-limiter/port"
)

func RateLimitMiddleware(rateLimiterService port.RateLimiterService, rateLimitAddressPerSecond int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		ctx := context.TODO()
		//apiToken := c.Request.Header["API_KEY"]

		err := rateLimiterService.Allow(ctx, ip, rateLimitAddressPerSecond)
		if err != nil {
			if errors.Is(err, ErrLimitExceeded) {
				c.String(http.StatusTooManyRequests, err.Error())
			} else {
				c.String(http.StatusInternalServerError, err.Error())
				log.Print("RateLimitMiddleware General Error -> ", err)
			}
			c.Abort()
			return
		}
		c.Next()
	}
}
