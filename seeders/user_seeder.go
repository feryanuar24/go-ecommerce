package seeders

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
	"log"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
)

func SeedUsers() {
	rand.Seed(time.Now().UnixNano())

	var users []models.User
	for i := 0; i < 100; i++ {
		user := models.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: "1234567890",
			Role:     randomRole(),
		}
		users = append(users, user)
	}

	for _, user := range users {
		err := config.DB.Create(&user).Error
		if err != nil {
			log.Printf("Failed to seed user %s: %v\n", user.Email, err)
		} else {
			log.Printf("User %s seeded successfully\n", user.Email)
		}
	}
}

func randomRole() string {
	roles := []string{"Admin", "Pengguna"}
	return roles[rand.Intn(len(roles))]
}
