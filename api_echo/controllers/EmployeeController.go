package controllers

import (
	"net/http"

	"github.com/belajar_golang/api_echo/models"

	"github.com/labstack/echo"
)

func GetEmployees(c echo.Context) error {
	result := models.GetEmployee()
	println("foo")
	return c.JSON(http.StatusOK, result)
}
