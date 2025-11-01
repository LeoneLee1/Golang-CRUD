package main

import (
	"go-crud/database"
	"go-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	r.GET("/api", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Golang CRUD REST API",
		})
	})

	routes.UserRoutes(r)

	r.Run(":8081")
}