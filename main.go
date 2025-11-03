package main

import (
	"go-crud/database"
	"go-crud/routes"

	_ "go-crud/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go CRUD REST API
// @version 1.0
// @description This is a CRUD API built with Golang + Gin.
// @termsOfService http://swagger.io/terms/

// @contact.name Daniel Lee
// @contact.email danielruntuwene42@gmail.com

// @host localhost:8081
// @BasePath /api
func main() {
	r := gin.Default()

	database.ConnectDatabase()

	r.GET("/api", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Golang CRUD REST API",
		})
	})

	
	routes.UserRoutes(r)
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8081")
}