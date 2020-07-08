package kerbalwzygo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"log"
	"net/http"
)

var ErrTooManyRequests = errors.New(http.StatusText(429))
var ErrTokenNumber = errors.New("secToken的值必须小于等于maxToken的值")

// 限流中间件
// 每秒产生secToken个令牌, 最多存放maxToken个令牌, 即一秒内最多允许maxToken个请求
func FlowLimitMiddleware(secToken rate.Limit, maxToken int) gin.HandlerFunc {
	if int(secToken) > maxToken {
		log.Fatal(ErrTokenNumber)
	}
	limiter := rate.NewLimiter(secToken, maxToken)
	return func(c *gin.Context) {
		if limiter.Allow() == false {
			_ = c.AbortWithError(429, ErrTooManyRequests)
		} else {
			c.Next()
		}
	}
}
