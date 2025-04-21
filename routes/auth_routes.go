package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"hotel_management_app_backend/controllers"
)

func AuthRoutes(r *gin.Engine) {

	userGroup := r.Group("/auth")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}
}

func LogoutRoute(r *gin.Engine) {
	store := cookie.NewStore([]byte("super-secret-key"))
	r.Use(sessions.Sessions("my-session", store)) // <- Important

	csrfGroup := r.Group("/")
	csrfGroup.Use(csrf.Middleware(csrf.Options{
		Secret: "very-secret-csrf-key",
		ErrorFunc: func(c *gin.Context) {
			c.JSON(403, gin.H{"error": "CSRF token invalid or missing"})
			c.Abort()
		},
	}))

	// Provide CSRF token to frontend (GET request)
	csrfGroup.GET("/csrf-token", func(c *gin.Context) {
		c.JSON(200, gin.H{"csrf_token": csrf.GetToken(c)})
	})

	// Protected logout
	csrfGroup.POST("/auth/logout", controllers.Logout)

}
