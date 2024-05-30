package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"naaga.me/booking-rest-api/db"
	"naaga.me/booking-rest-api/models"
)

func main() {

	// Initializing the SQLite Database
	db.InitDB()

	// Creating the Go Gin Server
	server := gin.Default()

	// Defining endpoints and the handler functions
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:eventId", getEvent)

	// Starting the long-running REST API server
	server.Run(":8080")

}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event requested"})
		return
	}
	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event requested"})
		return
	}
	context.JSON(http.StatusOK, event)
}

// context *gin.Context is always passed down to the function as long as it's defined as handler function in the server GET/POST methods
func getEvents(context *gin.Context) {
	// Calling the events model function to get all Events
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events data"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// Creating an empty event variable of type Event struct and binding the POST body JSON to it. The function will return an error
	// if Struct binding constraints are not met
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Malformed body request"})
		return
	}

	event.ID = 1
	event.UserId = 1

	// Saving the event and responding back the status to the user
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create the event", "event": event})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully!", "event": event})
}
