package main

import (
	"benchmark/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		data := models.Person{Name: "John", Surname: "Doe"}
		return c.JSON(http.StatusOK, data)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
