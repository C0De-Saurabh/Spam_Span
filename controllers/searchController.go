package controllers

import (
	"Spam_Span/services"
	"Spam_Span/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SearchPhone(c *gin.Context) {
	// Retrieve the phone number from the query parameters
	phone := c.DefaultQuery("phone", "")

	// Validate the phone number
	if phone == "" || !utils.IsValidPhone(phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing phone number"})
		return
	}

	// Get the logged-in user's ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Check if the phone number belongs to a registered user
	user, err := services.GetUserByPhone(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If a registered user is found, check if the searching user is in their contact list
	if user != nil {
		// Check if the user is in the contact list of the searched user
		inContactList, err := services.IsUserInContactList(userID.(uint), phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Prepare response based on whether the user is in the contact list
		contactInfo := gin.H{
			"name":            user.Name,
			"phone":           user.Phone,
			"spam_likelihood": services.GetSpamLikelihood(phone),
		}

		// Include email if the user is in the contact list
		if inContactList {
			contactInfo["email"] = user.Email
		}

		c.JSON(http.StatusOK, gin.H{
			"contacts": []gin.H{contactInfo},
		})
		return
	}

	// If no registered user is found, return global contacts
	globalContacts, err := services.GetGlobalContactsByPhone(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return global contacts matching the phone number
	c.JSON(http.StatusOK, gin.H{
		"contacts": globalContacts,
	})
}

// Define a struct to hold the contact and its spam likelihood
type ContactWithSpamLikelihood struct {
	Name           string  `json:"name"`
	Phone          string  `json:"phone"`
	Email          *string `json:"email,omitempty"` // Email is optional, will be displayed if the user is registered and in the contact's list
	SpamLikelihood float64 `json:"spam_likelihood"` // Calculated spam likelihood for the contact
}

func SearchName(c *gin.Context) {
	// Get the name from the query parameter
	name, exists := c.GetQuery("name")
	if !exists || strings.TrimSpace(name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	// Trim whitespace from the name
	name = strings.TrimSpace(name)

	// Validate name length and format using utils package
	if len(name) < 2 || !utils.IsValidName(name) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name must be at least 2 characters long and contain only valid characters"})
		return
	}

	// Call the service to search for contacts by name
	contacts, err := services.SearchName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No contacts found"})
		return
	}

	// Prepare a slice to hold the contacts with calculated spam likelihood
	var contactsWithSpamLikelihood []ContactWithSpamLikelihood

	for _, contact := range contacts {
		// Get the spam likelihood for the contact
		spamLikelihood := services.GetSpamLikelihood(contact.Phone)

		// Create a response structure
		contactWithSpam := ContactWithSpamLikelihood{
			Name:           contact.Name,
			Phone:          contact.Phone,
			SpamLikelihood: spamLikelihood,
		}

		// Append the contact with spam likelihood to the response slice
		contactsWithSpamLikelihood = append(contactsWithSpamLikelihood, contactWithSpam)
	}

	// Return the search results with spam likelihood and email if applicable
	c.JSON(http.StatusOK, gin.H{
		"contacts": contactsWithSpamLikelihood,
	})
}
