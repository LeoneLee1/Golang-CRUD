package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userController := controllers.UserController{}

	user := r.Group("/api/users")

	{
		user.GET("", userController.GetUsers)
		user.POST("/create", userController.CreateUsers)
		user.GET("/:id", userController.GetUsersByID)
		user.PUT("/update/:id", userController.UpdateUsers)
		user.DELETE("/delete/:id",userController.DeleteUsers)
	}
}