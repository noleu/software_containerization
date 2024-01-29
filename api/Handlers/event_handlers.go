package Handlers

import (
	"net/http"
	"software_containerization/DTOs"
	"software_containerization/Models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEventHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newEvent Models.Event

		// Bind the JSON data to newEvent
		if err := c.ShouldBindJSON(&newEvent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Create(&newEvent); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusCreated, newEvent)
	}
}

func GetEventsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventId := c.Query("id")

		if eventId != "" {
			var event Models.Event
			if result := db.First(&event, eventId); result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
				return
			}
			c.JSON(http.StatusOK, event)
			return
		}

		var events []Models.Event
		if result := db.Find(&events); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, events)
	}
}

func DeleteEventHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventId := c.Param("id")

		if result := db.Delete(&Models.Event{}, eventId); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
	}
}

func UpdateEventHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input DTOs.UpdateEventInput
		eventId := c.Param("id")

		// Bind the JSON data to input
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Find the event to be updated
		var event Models.Event
		if result := db.First(&event, eventId); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}

		// Update event in the database
		if result := db.Model(&event).Updates(input); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, event)
	}
}
