package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rl "github.com/sirius1b/go-rate-limit/pkg"
)

type RateLimitHandler func(c *gin.Context)

func GinRateLimitMiddleware(limiter rl.IRateLimiter, handler RateLimitHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP() // Or any other identifier you want to use
		if !limiter.Allow(key) {
			handler(c) // Call the custom handler
			c.Abort()  // Prevent further processing of the request
			return
		}

		c.Next() // Proceed to the next handler
	}
}

func DefaultRateLimitHandler(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
		"message": "Too many requests",
	})
}
