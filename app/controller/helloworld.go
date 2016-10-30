package controller

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"time"
)

func Hello(c echo.Context) error {

	randomErr()

	return c.String(http.StatusOK, "Hello, World!")
}

func randomErr() {
	rand.Seed(time.Now().Unix())
	if rand.Intn(2) == 0 {
		log.Error("log error randomly")
	}
}
