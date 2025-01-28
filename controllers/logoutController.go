package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutHandler(c *gin.Context) {
	// Logout response (client should discard the token)
	//TO BE HANDLED BY CLIENT SIDE
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
