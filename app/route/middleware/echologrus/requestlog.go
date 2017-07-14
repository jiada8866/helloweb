// Package echologrus provides a middleware for echo that logs request details
// via the logrus logging library
package echologrus

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

// New returns a new middleware handler with a default name and logger
func New() echo.MiddlewareFunc {
	return NewWithLogger(logrus.StandardLogger())
}

// Another variant for better performance.
// With single log entry.
func NewWithLogger(l *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			isError := false

			body := c.Request().Body
			dataIn, _ := ioutil.ReadAll(body)
			if len(dataIn) > 0 {
				c.Request().Body = ioutil.NopCloser(bytes.NewReader(dataIn))
			}

			if err := next(c); err != nil {
				c.Error(err)
				isError = true
			}

			latency := time.Since(start)

			host, _ := os.Hostname()

			// TODO 完善access日志的field
			entry := l.WithFields(logrus.Fields{
				"type":    "access",
				"server":  host,
				"method":  c.Request().Method,
				"ip":      c.Request().RemoteAddr,
				"status":  c.Response().Status,
				"latency": latency.Nanoseconds() / int64(time.Millisecond),
				"body":    string(dataIn),
				"route":   c.Path(),
			})

			if c.Response().Status != http.StatusNotFound {
				entry = entry.WithField("api", c.Request().URL)
			} else {
				entry = entry.WithField("illegal_api", c.Request().URL)
			}

			if reqID := c.Request().Header.Get("X-Request-Id"); reqID != "" {
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
