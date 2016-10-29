package controller

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"net/http"
)

func Hello(c echo.Context) error {

	//log.WithFields(log.Fields{
	//	"api": "Hell0",
	//	"index":   1,
	//}).Info("First API is working!")

	return c.String(http.StatusOK, "Hello, World!")
}
