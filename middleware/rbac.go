package middleware

import (
	"net/http"
	"strconv"
	"test_backend/database"
	"test_backend/models"

	"github.com/gin-gonic/gin"
)

func RBAC(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.GetHeader("X-User-ID")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
			c.Abort()
			return
		}

		var user models.User
		err = database.DB.Preload("Role.Permissions").First(&user, userID).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		for _, perm := range user.Role.Permissions {
			if perm.Name == permission {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}
