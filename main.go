package main

import (
	"fmt"
	"github.com/jiada8866/helloweb/app/route"
	"github.com/jiada8866/helloweb/app/route/middleware/echologrus"
	"github.com/jiada8866/helloweb/app/shared/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"os"
)

var logpath string = "/tmp/log/helloweb.log"

func main() {
	logfile, err := os.Create(logpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logfile.Close()
	logger.Init(logfile)

	e := echo.New()

	e.Use(echologrus.New())
	e.Use(middleware.Recover())

	route.AddRouters(e)

	e.Run(standard.New(":1323"))
}
