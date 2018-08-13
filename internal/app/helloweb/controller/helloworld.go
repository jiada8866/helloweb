package controller

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func One(c echo.Context) error {
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
