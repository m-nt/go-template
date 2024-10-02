package internal

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// DefaultStructuredLogger logs a gin HTTP request in JSON format. Uses the
// default logger from rs/zerolog.
func DefaultStructuredLogger() gin.HandlerFunc {
	return StructuredLogger(&log.Logger)
}

type UserAgent struct {
	Os        string
	OsVersion string
	Mobile    bool
}

// StructuredLogger logs a gin HTTP request in JSON format. Allows to set the
// logger for testing purposes.
func StructuredLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timer
		path := c.Request.URL.Path
		//raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		//hooked := logger.Hook(NewLokiLogger())
		hooked := logger

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.Path = path

		ua := useragent.Parse(c.Request.Header.Get("User-Agent"))
		// Log using the params
		var logEvent *zerolog.Event
		if c.Writer.Status() >= 500 {
			logEvent = hooked.Error()
		} else {
			logEvent = hooked.Info()
		}

		logEvent.Str("client_id", param.ClientIP).
			Str("method", param.Method).
			Int("status_code", param.StatusCode).
			Str("path", param.Path).
			Str("latency", param.Latency.String()).
			Interface("user_agent", UserAgent{
				Os:        ua.OS,
				Mobile:    ua.Mobile,
				OsVersion: ua.OSVersion,
			}).
			Msg(param.ErrorMessage)
	}
}
