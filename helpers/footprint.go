package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	defaultLogFormat            = "%s %-7s %s %s %3d %s [remote ip: \"%15s\", proto: \"%s\", agent: \"%s\"] -> \"%s\"(request id: %s) %v %13v error: \"%s\"\n"
	defaultLogFormatWithReqBody = "%s %-7s %s %s %3d %s [remote ip: \"%15s\", proto: \"%s\", agent: \"%s\"] -> \"%s\"(request id: %s) %v %13v request: \"%s\" error: \"%s\"\n"
)

func NewLoggerConfig(formatter func(gin.LogFormatterParams) string, skipPaths []string) gin.LoggerConfig {
	return gin.LoggerConfig{
		Formatter: formatter,
		SkipPaths: skipPaths,
	}
}

func DefaultLoggerFormatter(param gin.LogFormatterParams) string {
	if param.Method != http.MethodGet {
		reqBody, _ := ioutil.ReadAll(param.Request.Body)
		return loggingWithReqBodyLog(param, string(reqBody))
	}

	return fmt.Sprintf("Default logging")
}

func loggingWithReqBodyLog(param gin.LogFormatterParams, reqBody string) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	return fmt.Sprintf(defaultLogFormatWithReqBody,
		methodColor, param.Method, resetColor,
		statusColor, param.StatusCode, resetColor,
		param.ClientIP,
		param.Request.Proto,
		param.Request.UserAgent(),
		param.Path,
		param.Request.Header.Get("X-Request-Id"),
		param.TimeStamp.Format(time.RFC3339),
		param.Latency,
		reqBody,
		param.ErrorMessage,
	)
}
