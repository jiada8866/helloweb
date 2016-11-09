package route

import (
	"github.com/jiada8866/helloweb/app/controller"
	"github.com/labstack/echo"
)

func AddRouters(e *echo.Echo) {
	e.GET("/", controller.Hello)
	e.GET("/one", controller.One)
	e.GET("/two", controller.Two)

	e.POST("/post", controller.GetPostData)

}
