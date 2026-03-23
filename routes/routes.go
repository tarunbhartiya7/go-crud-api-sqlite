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
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
