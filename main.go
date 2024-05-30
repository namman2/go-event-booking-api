package main

import (
	"github.com/gin-gonic/gin"
	"naaga.me/booking-rest-api/db"
	"naaga.me/booking-rest-api/routes"
)

func main() {

	// Initializing the SQLite Database
	db.InitDB()

	// Creating the Go Gin Server
	server := gin.Default()

	// Defining endpoints and the handler functions
	routes.RegisterRoutes(server)

	// Starting the long-running REST API server
	server.Run(":8080")

}
