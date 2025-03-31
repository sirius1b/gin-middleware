package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirius1b/gin-middleware/pkg"

	rl "github.com/sirius1b/go-rate-limit/pkg"
)

func GetLimiter() *rl.IRateLimiter {

	options := rl.Options{
		Limit:  1,
		Window: time.Second,
	}

	limiter, err := rl.Require(rl.FixedWindow, options)

	if err != nil {
		fmt.Errorf("%s", err.Error())
	}

	return &limiter
}

func main() {
	router := gin.New()

	router.Use(pkg.ClientIPLimiter(
		GetLimiter(),
		pkg.DefaultRateLimitHandler,
	))

	router.GET("/test", func(c *gin.Context) {
		fmt.Printf("Request At: %v\n", time.Now())
	})

	router.Run(":8080")
}
