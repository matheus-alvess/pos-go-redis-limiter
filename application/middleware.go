package application

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pos-go-redis-limiter/infrastructure"
	"pos-go-redis-limiter/port"
)

func RateLimitMiddleware(rateLimiterService port.RateLimiterService, config *infrastructure.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ip        = c.ClientIP()
			ctx       = context.TODO()
			apiToken  = c.Request.Header.Get("API_KEY")
			limitTime = config.IpRateLimitPerSecond
			keyLock   = ip
		)

		if apiToken != "" {
			limitTime = config.TokenRateLimitPerSecond
			keyLock = fmt.Sprintf("%s:%s", ip, apiToken)
		}

		ok, err := rateLimiterService.Allow(ctx, keyLock, limitTime, config.GeneralTimeBan)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			log.Print("RateLimitMiddleware General Error -> ", err)
			c.Abort()
			return
		}

		if !ok {
			c.String(http.StatusTooManyRequests, "you have reached the maximum number of requests or actions allowed within a certain time frame")
			c.Abort()
			return
		}
		c.Next()
	}
}
