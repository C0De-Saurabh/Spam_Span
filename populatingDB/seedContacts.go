package populatingDB

import (
	"Spam_Span/config"
	"Spam_Span/models"
	"log"
	"time"

	"github.com/bxcodec/faker/v4"
)

func SeedContacts(userCount, contactCount int) {
	for i := 1; i <= userCount; i++ { // Assuming User IDs are sequential
		for j := 0; j < contactCount; j++ {
			contact := models.Contact{
				UserID:    uint(i),
				Name:      faker.Name(),
				Phone:     faker.Phonenumber(),
				Email:     faker.Email(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			// Add the contact to the database
			if err := config.DB.Create(&contact).Error; err != nil {
				log.Printf("Error seeding contact for user %d: %v", i, err)
			}
		}
	}
	log.Printf("%d contacts seeded successfully for each user", contactCount)
}
