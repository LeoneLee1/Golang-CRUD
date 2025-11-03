package routes

import (
	"go-crud/controllers"
	"go-crud/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userController := controllers.UserController{}
	authController := controllers.AuthController{}

	auth := r.Group("/api")

	// AUTH ROUTES
	{
		auth.POST("/api/login", authController.Login)
	}


	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())


	// USER ROUTES
	{
		protected.GET("/users", userController.GetUsers)
		protected.POST("/users/create", userController.CreateUsers)
		protected.GET("/users/:id", userController.GetUsersByID)
		protected.PUT("/users/update/:id", userController.UpdateUsers)
		protected.DELETE("/users/delete/:id",userController.DeleteUsers)
		protected.GET("/users/profile",userController.Profile)
	}
}