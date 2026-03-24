package routes

import (
	"log"
	"net/http"

	"example.com/events/models"
	"example.com/events/utils"
	"github.com/gin-gonic/gin"
)

// signUp godoc
// @Summary Register user
// @Description Create a new user account.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body signupRequest true "Signup payload"
// @Success 200 {object} signupResponse
// @Failure 400 {object} messageResponse
// @Failure 500 {object} messageResponse
// @Router /signup [post]
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

// login godoc
// @Summary Login user
// @Description Authenticate a user and return a JWT token.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body loginRequest true "Login payload"
// @Success 200 {object} loginResponse
// @Failure 400 {object} messageResponse
// @Failure 401 {object} messageResponse
// @Failure 500 {object} messageResponse
// @Router /login [post]
func login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
		return
	}
	err := user.ValidateCredentials()
	if err != nil {
		log.Printf("Login failed: %v", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
