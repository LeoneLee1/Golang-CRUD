package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/api/users")

	{
		user.GET("", controllers.GetUsers)
		user.POST("/create", controllers.CreateUsers)
		user.GET("/:id", controllers.ShowUsers)
		user.PUT("/update/:id", controllers.UpdateUsers)
		user.DELETE("/delete/:id", controllers.DeleteUsers)
	}
}