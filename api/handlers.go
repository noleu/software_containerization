package main

import (
	"net/http"
	"software_containerization/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRolesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roles []Models.Role
		if result := db.Find(&roles); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, roles)
	}
}

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser Models.User
		
		// Bind the JSON data to newUser
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add user to the database
		if result := db.Create(&newUser); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, newUser)
	}
}

func GetUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []Models.User
		if result := db.Find(&users); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}