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

// getEvents godoc
// @Summary List events
// @Description Get all events.
// @Tags events
// @Produce json
// @Success 200 {array} models.Event
// @Router /events [get]
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

// createEvent godoc
// @Summary Create event
// @Description Create a new event for the authenticated user.
// @Tags events
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body createEventRequest true "Event payload"
// @Success 201 {object} models.Event
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /events [post]
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

// getEventById godoc
// @Summary Get event
// @Description Get an event by ID.
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} models.Event
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /events/{id} [get]
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

// updateEvent godoc
// @Summary Update event
// @Description Update an event owned by the authenticated user.
// @Tags events
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Event ID"
// @Param request body updateEventRequest true "Event payload"
// @Success 200 {object} models.Event
// @Failure 400 {object} errorResponse
// @Failure 403 {object} messageResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /events/{id} [put]
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

// deleteEvent godoc
// @Summary Delete event
// @Description Delete an event owned by the authenticated user.
// @Tags events
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Event ID"
// @Success 200 {object} messageResponse
// @Failure 400 {object} errorResponse
// @Failure 403 {object} messageResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /events/{id} [delete]
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
