package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func Alert(c echo.Context) error {
	return c.JSON(http.StatusOK,"ok")
}
