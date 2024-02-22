package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

//GET
func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get events!"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

//POST
func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserID = userId

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save the event!"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully!",
		"event":   event,
	})
}

//GET by id
func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID!"})
		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get the event!"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

//PUT
func updateEvent(ctx *gin.Context){
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID!"})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get the event!"})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to update this event!"})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update the event!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

//DELETE
func deleteEvent(ctx *gin.Context) {
	eventId := ctx.Param("id")
	id, err := strconv.ParseInt(eventId, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID!"})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventByID(id)


	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete the event!"})
		return
	}

	if userId != event.UserID {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to delete this event!"})
		return
	}

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete the event!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}