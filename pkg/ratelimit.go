package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	rl "github.com/sirius1b/go-rate-limit/pkg"
)

type Handler func(c *gin.Context)

func ClientIPLimiter(limiter *rl.IRateLimiter, handler Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		if !(*limiter).Allow(key) {
			handler(c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequestHeaderLimiter(header string, limiter *rl.IRateLimiter, handler Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader(header)
		if !(*limiter).Allow(key) {
			handler(c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func DefaultRateLimitHandler(c *gin.Context) {
	fmt.Printf("Rate Limitted At: %v\n", time.Now())
	c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
		"message": "Too many requests",
	})
}
