package kerbalwzygo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

var CustomLogger *GinStyleLogger

// WARNING!!!
// This gin style logger can not output the colorful log information in terminal, only pure text.
// Usage:
func init() {
	CustomLogger = NewGinStyleLogger(nil, nil)
}

// Logger for output the custom log information with gin style in the gin.HandlerFunc.
type GinStyleLogger struct {
	// Optional. Default value is gin.defaultGinStyleLogFormatter
	Formatter gin.LogFormatter

	// Output is a writer where logs are written.
	// Optional. Default value is gin.DefaultWriter.
	Output io.Writer
}

// Output the log information with custom status code and message.
func (p *GinStyleLogger) Fprintln(c *gin.Context, statusCode int, message string) {
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	param := gin.LogFormatterParams{
		Request: c.Request,
		Keys:    c.Keys,
	}
	param.TimeStamp = time.Now()
	param.ClientIP = c.ClientIP()
	param.Method = c.Request.Method
	param.StatusCode = statusCode
	param.ErrorMessage = message

	if raw != "" {
		path = path + "?" + raw
	}
	param.Path = path
	_, _ = fmt.Fprintln(p.Output, p.Formatter(param))
}

// defaultGinStyleLogFormatter is the default log format function gin.StyleLogger uses.
var defaultGinStyleLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %s | %s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		"ginStyleLog",
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// return a StyleLogger instance with the params. when the param is nil, will use the default value.
func NewGinStyleLogger(out io.Writer, formatter gin.LogFormatter) *GinStyleLogger {
	if out == nil {
		out = gin.DefaultWriter
	}
	if formatter == nil {
		formatter = defaultGinStyleLogFormatter
	}
	return &GinStyleLogger{Output: out, Formatter: formatter}
}
