package services

import (
	"Spam_Span/config"
	"Spam_Span/models"
	"errors"
	"gorm.io/gorm"
)

// AddContact adds a contact for a specific user
func AddContact(userID uint, contactName string, contactPhone string) error {
	// Validate inputs
	if contactName == "" || contactPhone == "" {
		return errors.New("contact name and phone are required")
	}

	// Create the contact
	contact := models.Contact{
		UserID: userID,
		Name:   contactName,
		Phone:  contactPhone,
	}

	// Save the contact to the database
	if err := config.DB.Create(&contact).Error; err != nil {
		return errors.New("failed to add contact")
	}

	// Sync with the Global Contacts table
	var globalContact models.GlobalContact
	result := config.DB.Where("phone = ?", contactPhone).First(&globalContact)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Add new contact to Global Contacts
			newGlobalContact := models.GlobalContact{
				Phone:        contactPhone,
				Name:         contactName,
				SpamReported: 0,
			}
			if err := config.DB.Create(&newGlobalContact).Error; err != nil {
				return errors.New("failed to sync with global contacts")
			}
		} else {
			return errors.New("failed to query global contacts")
		}
	}

	return nil
}

// GetUserContacts fetches all personal contacts of the logged-in user
func GetUserContacts(userID uint) ([]models.Contact, error) {
	var contacts []models.Contact

	// Fetch contacts linked to the user
	if err := config.DB.Where("user_id = ?", userID).Find(&contacts).Error; err != nil {
		return nil, errors.New("failed to fetch contacts")
	}

	return contacts, nil
}
