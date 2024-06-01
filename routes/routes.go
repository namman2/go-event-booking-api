package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/signup", userSignup)
	server.POST("/login", login)
	server.GET("/events/:eventId", getEvent)

	// Routes protected by JWT
	server.POST("/events", createEvent)
	server.PUT("/events/:eventId", updateEvent)
	server.DELETE("/events/:eventId", deleteEvent)

}
