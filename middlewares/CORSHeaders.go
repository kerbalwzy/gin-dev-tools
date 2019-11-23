package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Add CORS headers in response message.
// Allow any origin and take the cookie data here, You may need change for your own demands.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", `Authorization, Content-Type, Content-Length, 
X-CSRF-Token, X-Requested-With, X-CustomHeader, Accept, Accept-Encoding, Accept-Language, Origin, Host, Connection, DNT, 
Keep-Alive, User-Agent, If-Modified-Since, Cache-Control, Pragma`)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, OPTIONS, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
