package Handlers

import (
	"net/http"
	"software_containerization/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"errors"
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

func GetRoleByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var role Models.Role
		roleId := c.Param("id") // Extracting the role ID from the URL

		if result := db.First(&role, roleId); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, role)
	}
}