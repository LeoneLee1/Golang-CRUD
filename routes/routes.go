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
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}


	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())


	// USER ROUTES
	{
		protected.GET("/users", userController.GetUsers)
		protected.GET("/users/:id", userController.GetUsersByID)
		protected.GET("/users/profile",userController.Profile)

		admin := protected.Group("/admin")
		admin.Use(middleware.AdminMiddleware())
		admin.POST("/users/create", userController.CreateUsers)
		admin.PUT("/users/update/:id", userController.UpdateUsers)
		admin.DELETE("/users/delete/:id",userController.DeleteUsers)
	}
}