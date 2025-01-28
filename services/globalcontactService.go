package services

import (
	"Spam_Span/config"
	"Spam_Span/models"
	"errors"
)

func AddToGlobalContacts(phone, name string) error {
	var globalContact models.GlobalContact
	result := config.DB.Where("phone = ?", phone).First(&globalContact)
	if result.RowsAffected > 0 {
		// Phone already exists in global contacts; no action needed
		return nil
	}

	// Add new entry to global contacts
	newContact := models.GlobalContact{
		Phone:        phone,
		Name:         name,
		SpamReported: 0,
	}

	if err := config.DB.Create(&newContact).Error; err != nil {
		return errors.New("failed to add contact to global database")
	}

	return nil
}

// GetGlobalContactsByPhone returns all global contacts matching the given phone number
func GetGlobalContactsByPhone(phone string) ([]models.GlobalContact, error) {
	var globalContacts []models.GlobalContact
	err := config.DB.Where("phone = ?", phone).Find(&globalContacts).Error
	if err != nil {
		return nil, err
	}
	return globalContacts, nil
}
