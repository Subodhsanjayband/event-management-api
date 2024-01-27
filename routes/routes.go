package routes

import "github.com/gin-gonic/gin"

func RegesterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents) //for fetching all the events

	server.GET("/events/:id", getEvent) // for fetching a specific event using id

	server.POST("/events", createEvent) // for creating new event

	server.PUT("/events/:id", updateEvent) // for updating a specific event by using id

	server.DELETE("/events/:id", deleteEvent) // for deleting the event by using id

	server.POST("/signup", signup)

	server.POST("/login", login)

	server.GET("/users", getUsers)

}
