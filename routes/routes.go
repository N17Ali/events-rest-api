package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/n17ali/events-rest-api/middlewares"
)

func ResigterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancleRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
