package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)

	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the events!"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event                        //aqui a variavel recebe o type Event
	err := context.ShouldBindBodyWithJSON(&event) //aqui garante que a variavel vai ter o body definido em models.Event

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the event!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created sucessfully!", "event": event})
}
