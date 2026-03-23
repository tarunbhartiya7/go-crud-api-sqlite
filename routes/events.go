package routes

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"example.com/events/models"
	"example.com/events/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		log.Printf("CreateEvent failed: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event data"})
		return
	}
	createdEvent := event.Save(userId)
	context.JSON(http.StatusCreated, createdEvent)
}

func getEventById(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		log.Printf("GetEventById failed: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
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
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEvent, err := models.UpdateEvent(eventId, event)
	if err != nil {
		log.Printf("UpdateEvent failed: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	err = models.DeleteEvent(eventId)
	if err != nil {
		log.Printf("DeleteEvent failed: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
