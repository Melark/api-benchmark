package main

import (
	"benchmark/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		data := models.Person{Name: "John", Surname: "Doe"}
		return c.JSON(data)
	})

	app.Listen(":5204")
}
