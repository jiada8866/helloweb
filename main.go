package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
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
	logger.Init(logfile, false)

	e := echo.New()

	e.SetDebug(true)

	// Logger as an io.Writer
	// 可能是最简单的用logrus打印echo自身日志的方式
	// 缺点：logrus打印的echo日志的level都是info，而且将echo日志都写在msg里
	w := log.StandardLogger().Writer()
	defer w.Close()
	e.SetLogOutput(w)

	e.Use(echologrus.New())
	e.Use(middleware.Recover())

	route.AddRouters(e)

	e.Run(standard.New(":1323"))
}
