package populatingDB

import (
	"Spam_Span/config"
	"Spam_Span/models"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
)

func SeedGlobalContacts(count int) {
	// Seed the random number generator for consistent randomness
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		globalContact := models.GlobalContact{
			Phone:        faker.Phonenumber(),
			Name:         faker.Name(),
			SpamReported: rand.Intn(11), // Random spam report count between 0 and 10
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		// Add the global contact to the database
		if err := config.DB.Create(&globalContact).Error; err != nil {
			log.Printf("Error seeding global contact: %v", err)
		}
	}
	log.Printf("%d global contacts seeded successfully", count)
}
