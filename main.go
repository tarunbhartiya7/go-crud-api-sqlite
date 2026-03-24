package main

import (
	"example.com/events/db"
	_ "example.com/events/docs"
	"example.com/events/routes"
	"github.com/gin-gonic/gin"
)

// @title Events CRUD API
// @version 1.0
// @description Event management API with auth and registrations.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	db.InitDB()
	server := gin.Default()
	routes.SetupRoutes(server)
	server.Run(":8080")
}
