package middleware

import (
	"github.com/jiadas/helloweb/pkg/shared/log"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func ContextLogus() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			reqID := req.Header.Get(echo.HeaderXRequestID)
			if reqID == "" {
				reqID = res.Header().Get(echo.HeaderXRequestID)
			}

			if reqID != "" {
				r := req.WithContext(log.WithLogger(req.Context(), logrus.WithField("request_id", reqID)))
				c.SetRequest(r)
			}

			return next(c)
		}
	}
}
