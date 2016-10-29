// Package echologrus provides a middleware for echo that logs request details
// via the logrus logging library
package echologrus

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"os"
)

// New returns a new middleware handler with a default name and logger
func New() echo.MiddlewareFunc {
	return NewWithName("web")
}

// NewWithName returns a new middleware handler with the specified name
func NewWithName(name string) echo.MiddlewareFunc {
	return NewWithNameAndLogger(name, logrus.StandardLogger())
}

// NewWithNameAndLogger returns a new middleware handler with the specified name
// and logger
func NewWithNameAndLogger(name string, l *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			host, _ := os.Hostname()

			entry := l.WithFields(logrus.Fields{
				"server":  host,
				"request": c.Request().URI(),
				"method":  c.Request().Method(),
				"remote":  c.Request().RemoteAddress(),
			})

			if reqID := c.Request().Header().Get("X-Request-Id"); reqID != "" {
				entry = entry.WithField("request_id", reqID)
			}

			entry.Info("started handling request")

			if err := next(c); err != nil {
				c.Error(err)
			}

			latency := time.Since(start)

			entry.WithFields(logrus.Fields{
				"status":      c.Response().Status(),
				"text_status": http.StatusText(c.Response().Status()),
				"took":        latency,
				fmt.Sprintf("measure#%s.latency", name): latency.Nanoseconds(),
			}).Info("completed handling request")

			return nil
		}
	}
}

// Another variant for better performance.
// With single log entry and time format.
func LogrusLogger(name string, l *logrus.Logger, timeFormat string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			isError := false

			if err := next(c); err != nil {
				c.Error(err)
				isError = true
			}

			latency := time.Since(start)

			entry := l.WithFields(logrus.Fields{
				"path":    c.Request().URI(),
				"method":  c.Request().Method(),
				"ip":      c.Request().RemoteAddress(),
				"status":  c.Response().Status(),
				"latency": latency,
				"time":    time.Now().Format(timeFormat),
			})

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
