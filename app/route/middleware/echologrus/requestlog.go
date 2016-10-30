// Package echologrus provides a middleware for echo that logs request details
// via the logrus logging library
package echologrus

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

// New returns a new middleware handler with a default name and logger
func New() echo.MiddlewareFunc {
	return NewWithLogger(logrus.StandardLogger())
}

// Another variant for better performance.
// With single log entry and time format.
func NewWithLogger(l *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			isError := false

			if err := next(c); err != nil {
				c.Error(err)
				isError = true
			}

			latency := time.Since(start)

			host, _ := os.Hostname()

			entry := l.WithFields(logrus.Fields{
				"type":    "access",
				"server":  host,
				"method":  c.Request().Method(),
				"ip":      c.Request().RemoteAddress(),
				"status":  c.Response().Status(),
				"latency": latency.Nanoseconds() / int64(time.Millisecond),
			})

			if c.Response().Status() != http.StatusNotFound {
				entry = entry.WithField("api", c.Request().URI())
			} else {
				entry = entry.WithField("illegal_api", c.Request().URI())
			}

			if reqID := c.Request().Header().Get("X-Request-Id"); reqID != "" {
				entry = entry.WithField("request_id", reqID)
			}

			// Check middleware error
			if isError {
				entry.Error("error by handling request")
			} else {
				entry.Info("completed handling request")
			}

			return nil
		}
	}
}
