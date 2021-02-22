package router

import (
	"fmt"
	"github.com/natae/golang/controller"

	"github.com/labstack/echo/v4"
)

func SetUserRoute(path string, echo *echo.Echo) {
	echo.POST(path + "/login", controller.PostLogin, UserRouteMiddleware)
}

func UserRouteMiddleware (next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("route middleware before")
		if err := next(c); err != nil {
			c.Error(err)
		}
		fmt.Println("route middleware after")

		ok := c.Get("OK").(string)
		fmt.Println("OK:", ok)

		return nil
	}
}
