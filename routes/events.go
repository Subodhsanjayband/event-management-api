package routes

import (
	"net/http"
	"strconv"

	"github.com/Subodhsanjayband/event_manager/models"
	"github.com/Subodhsanjayband/event_manager/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	/*message := gin.H{
		"message": "Hi there bitch!!!!",
	}*/
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could fetch events.",
		})
		return
	}
	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could fetch event id: " + context.Param("id"),
		})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could fetch event id: " + context.Param("id"),
		})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
		})
	}

	id, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Not authorised",
		})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	event.UserID = id
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save the events,please try later.",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created...", "event": event,
	})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could fetch event id: " + context.Param("id"),
		})
		return
	}
	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could fetch event id: " + context.Param("id"),
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}
	updatedEvent.ID = eventId
	event, err := updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to update event " + context.Param("id"),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated the event", "updated_event": event})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could fetch event id: " + context.Param("id"),
		})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "The event does not exist for id: " + context.Param("id"),
		})
		return
	}
	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to delete the event for id: " + context.Param("id"),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted the record", "event": event,
	})
}
