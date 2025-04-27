package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Protected(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized!"})
}
