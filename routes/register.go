package routes

import (
	"errors"
	"net/http"
	"strconv"

	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt("userId")
	id := context.Param("id")
	eventId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event"})
}

func cancelRegistrationForEvent(context *gin.Context) {
	userId := context.GetInt("userId")
	id := context.Param("id")
	eventId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if errors.Is(err, models.ErrRegistrationNotFound) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Registration not found"})
		return
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled"})
}
