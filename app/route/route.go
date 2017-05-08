package route

import (
	"fmt"

	"github.com/jiadas/helloweb/app/controller"
	"github.com/labstack/echo"
)

func AddRouters(e *echo.Echo) {
	e.GET("/", controller.Hello)
	e.GET("/one", controller.One)
	e.GET("/two", controller.Two)

	e.POST("/post", controller.GetPostData)

	e.POST("/alert", controller.Alert)

	// test echo group
	groupTest(e)
}

// groupTest测试echo的分组路由功能
func groupTest(e *echo.Echo) {
	// 有共同的前缀"group"
	g := e.Group("/group")
	g.GET("/one", func(c echo.Context) error {
		return c.JSON(200, "g1")
	})
	g.GET("/two", func(c echo.Context) error {
		return c.JSON(200, "g2")
	})
	g.GET("/three", func(c echo.Context) error {
		return c.JSON(200, "g3")
	})

	one := e.Group("/one")
	one.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 增加该middleware为了验证，新建的分组one，不包括之前的"/one"
			// 即 用one分组指定的路由才会用到one分组use的middleware
			fmt.Println("come into group 'one'")
			return next(c)
		}
	})
	one.GET("/g0", func(c echo.Context) error {
		return c.JSON(200, "/one/g0 in group 'one'")
	})

	// e.GET("/one/g1")存在的话，该路由永远进不去
	one.GET("/g1", func(c echo.Context) error {
		return c.JSON(200, "/one/g1 in group 'one'")
	})

	// 同时存在2个"/one/g1"，若请求"/one/g1"，进的是这个路由
	e.GET("/one/g1", func(c echo.Context) error {
		return c.JSON(200, "/one/g1 without group")
	})
}
