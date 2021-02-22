package main

import (
	"github.com/natae/golang/router"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

func main() {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			body, err := ioutil.ReadAll(c.Request().Body)
			if err != nil {
				fmt.Println("err", err)
			}
			fmt.Println("body:", body)
			c.Set("body", body)

			fmt.Println("middleware 1 start")
			if err := next(c); err != nil {
				c.Error(err)
			}
			fmt.Println("middleware 1 end")
			return nil
		}
	})

	router.SetUserRoute("/user", e)

	e.Logger.Fatal(e.Start(":5001"))
}

