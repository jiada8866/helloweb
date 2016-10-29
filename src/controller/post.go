package controller

import (
	"github.com/labstack/echo"
)

type TestPost struct {
	Images  []string `json:"images"`
	Content string   `json:"content"`
}

func GetPostData(c echo.Context) error {
	tp := new(TestPost)
	if err := c.Bind(tp); err != nil {
		return err
	}

	return c.JSON(200, tp)
}
