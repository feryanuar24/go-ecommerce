package seeders

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
	"log"
)

func SeedProductCategory() {
	product_categories := []models.ProductCategory{
		{
			Name: "Makanan",
		},
		{
			Name: "Minuman",
		},
		{
			Name: "Pakaian",
		},
		{
			Name: "Elektronik",
		},
		{
			Name: "Kesehatan",
		},
		{
			Name: "Olahraga",
		},
		{
			Name: "Kecantikan",
		},
		{
			Name: "Perabotan",
		},
		{
			Name: "Mainan",
		},
		{
			Name: "Buku",
		},
	}

	for _, product_category := range product_categories {
		err := config.DB.Create(&product_category).Error
		if err != nil {
			log.Printf("Failed to seed Product Category %s: %v\n", product_category.Name, err)
		} else {
			log.Printf("Porduct Category %s seeded successfully\n", product_category.Name)
		}
	}
}
