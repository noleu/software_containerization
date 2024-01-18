package Handlers

import (
	"net/http"
	"software_containerization/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser Models.User
		// Bind the JSON data to newUser
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		newUser.Password = string(hashedPassword)

		// Add user to the database
		if result := db.Create(&newUser); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusCreated, newUser)
	}
}

func GetUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []Models.User
		if result := db.Find(&users); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func GetUserByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Models.User
		userId := c.Param("id") // Extracting the user ID from the URL

		if result := db.First(&user, userId); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func GetUserByEmailHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Models.User
		userEmail := c.Param("email") // Extract the email from the URL

		if result := db.Where("email = ?", userEmail).First(&user); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input Models.UpdateUserInput
		userId := c.Param("id") // Extracting the user ID from the URL

		// Bind the JSON data to input
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Find and update the user
		var user Models.User
		if result := db.First(&user, userId); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// If a new password is provided, hash it
		if input.Password != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
				return
			}
			input.Password = string(hashedPassword)
		} else {
			// Exclude the Password field from update if it's empty
			db = db.Omit("Password")
		}

		db.Model(&user).Updates(input)
		c.JSON(http.StatusOK, user)
	}
}