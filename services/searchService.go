package services

import (
	"Spam_Span/config"
	"Spam_Span/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// GetUserByPhone checks if the phone number belongs to a registered user
func GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		// If no registered user is found, return nil
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// IsUserInContactList checks if a given user is in the contact list of another user
func IsUserInContactList(userID uint, contactPhone string) (bool, error) {
	var contact models.Contact
	err := config.DB.Where("user_id = ? AND phone = ?", userID, contactPhone).First(&contact).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not in contact list
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// SearchName retrieves global contacts based on the name search query
func SearchName(name string) ([]models.GlobalContact, error) {
	var globalContacts []models.GlobalContact

	// Search for global contacts matching the name, with or without email if available
	err := config.DB.
		Preload("User").                    // Preload the associated user data to get the email
		Where("name LIKE ?", "%"+name+"%"). // Search for the name with wildcard
		Find(&globalContacts).Error         // Find the matching global contacts

	if err != nil {
		return nil, fmt.Errorf("failed to search contacts: %v", err)
	}

	return globalContacts, nil
}
