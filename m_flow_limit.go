package kerbalwzygo

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"errors"
	"net/http"
)

var ErrTooManyRequests = errors.New(http.StatusText(429))

var limiter = rate.NewLimiter(100, 500)

// Flow limit middleware, only allow 500 request in max at the same one second here.
// The core principle is token bucket
func FlowLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.Allow() == false {
			_ = c.AbortWithError(429, ErrTooManyRequests)
			return
		}

		c.Next()
	}
}
