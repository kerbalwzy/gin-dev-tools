package kerbalwzygo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域处理中间
func CORSMiddleware(origin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		SetCORSHeaders(c.Writer.Header(), origin)
		if c.Request.Method == "OPTIONS" || c.Request.Method == "HEAD" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func SetCORSHeaders(header http.Header, allowOrigin string) {
	header.Set("Access-Control-Allow-Origin", allowOrigin)
	header.Set("Access-Control-Allow-Credentials", "true")
	header.Set("Access-Control-Allow-Headers", `Authorization, Content-Type, Content-Length, 
X-CSRF-Token, X-Requested-With, X-CustomHeader, Accept, Accept-Encoding, Accept-Language, Origin, Host, Connection, DNT, 
Keep-Alive, User-Agent, If-Modified-Since, Cache-Control, Pragma`)
	header.Set("Access-Control-Allow-Methods", "HEAD, GET, OPTIONS, POST, PUT, DELETE")
	header.Set("Access-Control-Max-Age", "86400")
}
