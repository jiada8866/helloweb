package route

import (
	"github.com/labstack/echo"

	"github.com/jiadas/helloweb/internal/app/helloweb/controller"
)

func AddRouters(e *echo.Echo) {
	e.GET("/", controller.Hello)
	e.GET("/one/:id", controller.One)
	e.GET("/two", controller.Two)

	e.POST("/post", controller.GetPostData)

	e.POST("/alert", controller.Alert)
}
