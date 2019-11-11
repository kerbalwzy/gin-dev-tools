package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"errors"
	"net/http"
)

var ErrTooManyRequests = errors.New(http.StatusText(429))

var limiter = rate.NewLimiter(100, 500)

// Control flow by Token bucket.
func FlowLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.Allow() == false {
			_ = c.AbortWithError(429, ErrTooManyRequests)
			return
		}
		c.Next()
	}
}
