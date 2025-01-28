package controllers

import (
	"Spam_Span/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddContactHandler adds a new contact for the logged-in user
func AddContactHandler(c *gin.Context) {
	var request struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	// Parse JSON request
	if err := c.ShouldBindJSON(&request); err != nil || request.Name == "" || request.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Add contact for the user
	if err := services.AddContact(userID.(uint), request.Name, request.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact added successfully"})
}

// ShowContactsHandler returns all personal contacts of the logged-in user
func ShowContactsHandler(c *gin.Context) {
	// Retrieve user ID from the context (set during authentication)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Fetch user contacts using the service function
	contacts, err := services.GetUserContacts(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"contacts": contacts})
}
