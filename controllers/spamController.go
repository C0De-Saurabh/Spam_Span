package controllers

import (
	"Spam_Span/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MarkAsSpamHandler(c *gin.Context) {
	var request struct {
		Phone string `json:"phone"`
	}

	// Parse JSON input
	if err := c.ShouldBindJSON(&request); err != nil || request.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone input"})
		return
	}

	// Get the user ID from the JWT or context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert the userID to uint
	uid, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the service to mark the phone as spam
	if err := services.MarkAsSpam(uid, request.Phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Marked as spam successfully"})
}
