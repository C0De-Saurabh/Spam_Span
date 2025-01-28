package populatingDB

import (
	"log"
	"time"

	"Spam_Span/config"
	"Spam_Span/models"
	"github.com/bxcodec/faker/v4"
)

func SeedUsers(count int) {
	for i := 0; i < count; i++ {
		user := models.User{
			Name:               faker.Name(),
			Phone:              faker.Phonenumber(),
			Password:           faker.Password(),
			Email:              faker.Email(),
			SpamReportsToday:   0,
			LastSpamReportDate: "",
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		// Add the user to the database
		if err := config.DB.Create(&user).Error; err != nil {
			log.Printf("Error seeding user: %v", err)
		}
	}
	log.Printf("%d users seeded successfully", count)
}
