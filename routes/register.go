package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/n17ali/events-rest-api/models"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func cancleRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.CancleRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not cancle registeration"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "cancled registration"})
}
