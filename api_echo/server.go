package main

import (
	"net/http"

	"github.com/belajar-golang/api_echo/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Welcome mvc echo with mysql app using Golang")
	})
	e.GET("/employees", controllers.GetEmployees)

	e.Logger.Fatal(e.Start(":8081"))

}
