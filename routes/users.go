package routes

import (
	"log"
	"net/http"

	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
		return
	}
	user, err := user.Save()
	if err != nil {
		log.Printf("SignUp failed: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    gin.H{"id": user.ID, "email": user.Email},
	})
}
