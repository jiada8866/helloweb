package main

import (
	"fmt"

	"github.com/NYTimes/logrotate"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/jiadas/helloweb/internal/app/helloweb/route"
	"github.com/jiadas/helloweb/pkg/shared/logger"
	"github.com/jiadas/helloweb/pkg/shared/middleware/echologrus"
)

var logpath string = "/tmp/log/helloweb.log"

func main() {
	// use logrotate.NewFile when log rated by logrotate
	logfile, err := logrotate.NewFile(logpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.Init(logfile, false)

	e := echo.New()

	// Logger as an io.Writer
	// 可能是最简单的用logrus打印echo自身日志的方式
	// 缺点：logrus打印的echo日志的level都是info，而且将echo日志都写在msg里
	w := log.StandardLogger().Writer()
	defer w.Close()
	e.Logger.SetOutput(w)

	// Elog实现了echo/log.Logger接口，可以将上述缺点很好解决
	//e.SetLogger(&(logger.Elog{log.New()}))
	//e.SetLogOutput(logfile)

	e.Logger = &(logger.Elog{Logger: log.New()})
	//e.Logger.SetLevel()

	e.Use(echologrus.New())
	e.Use(middleware.Recover())

	route.AddRouters(e)

	log.WithFields(log.Fields{
		"type": "start",
		"addr": ":1323",
	}).Info("server is running")
	e.Start(":25378")
}
