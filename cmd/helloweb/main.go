package main

import (
	"fmt"

	"github.com/NYTimes/logrotate"
	"github.com/jiadas/helloweb/internal/app/helloweb/route"
	"github.com/jiadas/helloweb/pkg/shared/logger"
	"github.com/jiadas/helloweb/pkg/shared/middleware/echologrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

var logpath = "/tmp/log/helloweb.log"

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
	// 可能是最简单的用 logrus 打印 echo 自身日志的方式
	// 缺点：logrus 打印的 echo 日志的 level 都是 info，而且将 echo 日志的内容都写在 msg 里
	// w := log.StandardLogger().Writer()
	// defer w.Close()
	// e.Logger.SetOutput(w)

	// Elog 实现了 echo/log.Logger 接口，可以将上述缺点很好解决
	e.Logger = &logger.Elog{Logger: log.StandardLogger()}

	e.Use(echologrus.New())
	e.Use(middleware.Recover())

	route.AddRouters(e)

	log.WithFields(log.Fields{"type": "start", "addr": ":25378"}).Info("server is running")

	e.Start(":25378")
}
