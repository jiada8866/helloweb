package main

import (
	"fmt"
	"github.com/jiada8866/echo-logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"os"
	"route"
	"shared/logger"
)

var logpath string = "/tmp/log/helloweb.log"
var timeFormat string = "2006-01-02 15:04:05"

func main() {
	logfile, err := os.Create(logpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logfile.Close()
	logger.Init(logfile)

	e := echo.New()

	e.Use(echologrus.NewWithTimeFormat(timeFormat))
	e.Use(middleware.Recover())

	route.AddRouters(e)

	e.Run(standard.New(":1323"))
}
