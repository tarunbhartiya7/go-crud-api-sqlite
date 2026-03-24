package routes

import (
	"example.com/events/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", middlewares.AuthMiddleware, createEvent)
	server.GET("/events/:id", getEventById)
	server.PUT("/events/:id", middlewares.AuthMiddleware, updateEvent)
	server.DELETE("/events/:id", middlewares.AuthMiddleware, deleteEvent)
	server.POST("/events/:id/register", middlewares.AuthMiddleware, registerForEvent)
	server.DELETE("/events/:id/register", middlewares.AuthMiddleware, cancelRegistrationForEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
