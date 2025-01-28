package main

import (
	"Spam_Span/config"
	"Spam_Span/controllers"
	"Spam_Span/middlewares"
	"Spam_Span/models"
	"Spam_Span/populatingDB"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initialize the database
func init() {
	var err error
	config.DB, err = gorm.Open(sqlite.Open("spam_span.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Migrate models
	config.DB.AutoMigrate(&models.User{}, &models.Contact{}, &models.GlobalContact{})
}

func main() {

	// Create a new Gin instance(ServeMux)
	r := gin.Default()

	// Seed data
	populatingDB.SeedUsers(10)          // Seed 10 users
	populatingDB.SeedContacts(10, 5)    // Seed 5 contacts per user
	populatingDB.SeedGlobalContacts(20) // Seed 20 global contacts

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authGroup := r.Group("/")
	authGroup.Use(middlewares.AuthMiddleware()) // Apply middleware to the group
	// Authenticated routes
	{
		authGroup.POST("/mark-spam", controllers.MarkAsSpamHandler)
		authGroup.GET("/search-phone", controllers.SearchPhone)
		authGroup.GET("/search-name", controllers.SearchName)
		authGroup.GET("/contacts", controllers.ShowContactsHandler)
		authGroup.POST("/logout", controllers.LogoutHandler) //TO BE IMPLEMENTED IN FRONTEND
		authGroup.POST("/add-contact", controllers.AddContactHandler)
	}

	// Start the server
	r.Run(":8080")
}
