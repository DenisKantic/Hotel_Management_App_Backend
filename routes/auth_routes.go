package routes

import (
	"github.com/gin-gonic/gin"
	"hotel_management_app_backend/controllers"
)

func AuthRoutes(r *gin.Engine) {
	userGroup := r.Group("/auth")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}
}
