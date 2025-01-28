package controllers

import (
	"Spam_Span/models"
	"Spam_Span/services"
	"Spam_Span/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User

	// Bind JSON input to user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate phone number
	if !utils.IsValidPhone(user.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	}

	// Hash the password
	user.Password = utils.HashPassword(user.Password)

	// If email is provided, validate it
	if user.Email != "" && !utils.IsValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Create user in DB
	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user - Already exists"})
		return
	}

	// Sync the registered user's data with Global Contacts
	if err := services.AddToGlobalContacts(user.Phone, user.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User registered but failed to sync with global contacts"})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var loginDetails struct {
		PhoneOrName string `json:"phone_or_name"`
		Password    string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if user exists
	user, err := services.GetUserByPhoneOrName(loginDetails.PhoneOrName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if !utils.CheckPasswordHash(loginDetails.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
