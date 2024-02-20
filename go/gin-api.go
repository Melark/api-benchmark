package main

import (
	"benchmark/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		data := models.Person{Name: "John", Surname: "Doe"}

		c.JSON(200, data)
	})
	r.Run(":5204") // listen and serve on 0.0.0.0:8080
}
