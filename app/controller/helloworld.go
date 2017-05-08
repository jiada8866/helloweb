package controller

import (
	"math/rand"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {

	randomErr()

	return c.String(http.StatusOK, "Hello, World!")
}

func One(c echo.Context) error {

	randomErr()

	return c.String(http.StatusOK, "One")
}

func Two(c echo.Context) error {

	randomErr()

	return c.String(http.StatusOK, "Two")
}

func randomErr() {
	rand.Seed(time.Now().Unix())
	if rand.Intn(2) == 0 {
		log.Error("log error randomly")
	}
}
