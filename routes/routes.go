package routes

import (
	"go-crud/controllers"
	"go-crud/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userController := controllers.UserController{}
	authController := controllers.AuthController{}

	r.POST("/api/login", authController.Login)

	user := r.Group("/api/users")
	user.Use(middleware.AuthMiddleware())

	{
		user.GET("", userController.GetUsers)
		user.POST("/create", userController.CreateUsers)
		user.GET("/:id", userController.GetUsersByID)
		user.PUT("/update/:id", userController.UpdateUsers)
		user.DELETE("/delete/:id",userController.DeleteUsers)
		user.GET("/profile",userController.Profile)
	}
}