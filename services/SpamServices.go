package services

import (
	"Spam_Span/config"
	"Spam_Span/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func MarkAsSpam(userID uint, phone string) error {
	// Fetch the user
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	// Reset the spam count if the date has changed
	if err := ResetSpamReportCount(&user); err != nil {
		return err
	}

	// Check if the user has reached the daily spam limit
	if user.SpamReportsToday >= 5 {
		return errors.New("you have already marked 5 numbers as spam today")
	}

	// Check if the phone exists in the global contacts table
	var globalContact models.GlobalContact
	result := config.DB.Where("phone = ?", phone).First(&globalContact)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Phone number not found; create a new entry in the global contacts table
			newContact := models.GlobalContact{
				Phone:        phone,
				Name:         "Unknown", // Default name for new entries
				SpamReported: 1,         // Initialize spam count as 1
			}
			if err := config.DB.Create(&newContact).Error; err != nil {
				return errors.New("failed to add phone number to global contacts")
			}
		} else {
			// Unexpected error
			return errors.New("failed to query global contacts")
		}
	} else {
		// Phone number exists; increment the spam count
		globalContact.SpamReported++
		if err := config.DB.Save(&globalContact).Error; err != nil {
			return errors.New("failed to update spam count in global contacts")
		}
	}

	// Increment the user's spam count
	user.SpamReportsToday++
	if err := config.DB.Save(&user).Error; err != nil {
		return errors.New("failed to update user's spam report count")
	}

	return nil
}

// ResetSpamReportCount resets the spam report count if the date has changed
func ResetSpamReportCount(user *models.User) error {
	currentDate := time.Now().Format("2006-01-02") // Current date in YYYY-MM-DD format

	// Reset the spam report count if the date has changed
	if user.LastSpamReportDate != currentDate {
		user.SpamReportsToday = 0
		user.LastSpamReportDate = currentDate
		if err := config.DB.Save(user).Error; err != nil {
			return errors.New("failed to reset spam report count")
		}
	}
	return nil
}

// GetSpamLikelihood calculates the spam likelihood for a given phone number
func GetSpamLikelihood(phone string) float64 {
	var phoneSpamReports int64
	var totalSpamReports int64

	// Fetch the total spam reports in the global contacts
	err := config.DB.Model(&models.GlobalContact{}).Where("spam_reported > 0").Count(&totalSpamReports).Error
	if err != nil {
		// If error occurs, log it and return a default value (0)
		fmt.Printf("Error fetching total spam reports: %v\n", err)
		return 0
	}

	// Fetch the spam reports for the given phone number
	err = config.DB.Model(&models.GlobalContact{}).Where("phone = ? AND spam_reported > 0", phone).Count(&phoneSpamReports).Error
	if err != nil {
		// If error occurs, log it and return a default value (0)
		fmt.Printf("Error fetching spam reports for phone number: %v\n", err)
		return 0
	}

	// If no total spam reports, return 0 to avoid division by zero
	if totalSpamReports == 0 {
		return 0
	}

	// Calculate the spam likelihood
	spamLikelihood := float64(phoneSpamReports) / float64(totalSpamReports) * 100

	return spamLikelihood
}
