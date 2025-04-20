package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hotel_management_app_backend/config"
	"hotel_management_app_backend/database"
	"hotel_management_app_backend/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cors_config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:5173"},
	}

	// making changes for pushing dev branch
	r := gin.Default()
	r.Use(cors.New(cors_config))
	config.LoadEnv()
	database.ConnectDB()
	fmt.Println("Hello World")

	// ROUTES
	routes.AuthRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		return
	}

}
