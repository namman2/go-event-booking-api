package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"naaga.me/booking-rest-api/models"
)

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

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	_, err = models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event requested for update"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Malformed body request for updating event"})
		return
	}
	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event requested for update"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	var deleteEvent models.Event
	deleteEvent.ID = id

	_, err = models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event requested for delete"})
		return
	}
	err = deleteEvent.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event requested for deletion"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully deleted the event"})
}
