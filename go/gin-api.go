package main

import (
	"benchmark/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		data := models.Person{Name: "John", Surname: "Doe"}

		c.JSON(200, data)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
