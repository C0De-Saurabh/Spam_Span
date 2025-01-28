package services

import (
	"Spam_Span/config"
	"Spam_Span/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUserByPhoneOrName(phoneOrName string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("phone = ? OR name = ?", phoneOrName, phoneOrName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
