package routes

import (
	"event-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// localhost:8080/events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	auth := server.Group("/")
	auth.Use(middlewares.Authenticate)
	auth.POST("/events", createEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)
	auth.POST("/events/:id/register", registerForEvent)
	auth.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}