package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"hotel_management_app_backend/database"
	"hotel_management_app_backend/models"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.UserEmployee

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please fill all fields"})
		fmt.Printf(err.Error())
		return
	}

	var existing models.UserEmployee
	if err := database.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email already exist"})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate password"})
		return
	}

	user.Password = string(hashed)

	// setting default role if not provided
	if user.Role == "" {
		user.Role = "worker"
	}

	// save user
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User registered successfully"})
}
