package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})
	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ini halaman users")
	})
	e.GET("/class", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "halaman dari /class")
	})
	e.Start(":8000")
}
