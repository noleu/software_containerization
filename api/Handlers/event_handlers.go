package Handlers

import (
	"net/http"
	"software_containerization/Models"
	"software_containerization/DTOs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"errors"
	"log"
)

func CreateEventHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newEvent Models.Event

		// Bind the JSON data to newEvent
		if err := c.ShouldBindJSON(&newEvent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate that the Organizer exists
		var organizer Models.User
		if err := db.First(&organizer, newEvent.OrganizerId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Organizer not found"})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		// Create the event if the organizer is valid
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
		userId := c.Query("userId")

		if eventId != "" {
			var event Models.Event
			if result := db.First(&event, eventId); result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
				return
			}
			c.JSON(http.StatusOK, event)
			return
		}

		if userId != "" {
			var events []Models.Event
			if result := db.Where("organizer_id = ?", userId).Find(&events); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
			c.JSON(http.StatusOK, events)
			return
		}

		var events []Models.Event
		if result := db.Preload("Organizer").Find(&events); result.Error != nil {
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

		log.Println("Hellloooooooooooooooooooooooo")
		log.Println(c.Param("userID"))
		
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