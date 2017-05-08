package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func Alert(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
