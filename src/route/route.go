package route

import (
	"controller"
	"github.com/labstack/echo"
)

func AddRouters(e *echo.Echo) {
	e.GET("/", controller.Hello)

	e.POST("/post", controller.GetPostData)

}
