package routes

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		log.Printf("CreateEvent failed: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event data"})
		return
	}
	userId, ok := context.Get("userId")
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	event.UserId = userId.(int)
	event.Save()
	context.JSON(http.StatusCreated, event)
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

	userId := context.GetInt("userId")
	log.Printf("Event user ID: %d, User ID: %d", event.UserId, userId)
	if event.UserId != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
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
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	userId := context.GetInt("userId")
	log.Printf("Event user ID: %d, User ID: %d", event.UserId, userId)
	if event.UserId != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
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
