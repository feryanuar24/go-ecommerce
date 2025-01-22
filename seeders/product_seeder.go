package seeders

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
	"log"
)

func SeedProduct() {
	products := []models.Product{
		{
			Name:              "Indomie Goreng",
			Description:       "Indomie Goreng",
			Price:             2500,
			ProductCategoryID: 1,
		},
		{
			Name:              "Teh Botol",
			Description:       "Teh Botol",
			Price:             5000,
			ProductCategoryID: 2,
		},
		{
			Name:              "Kaos Polos",
			Description:       "Kaos Polos",
			Price:             30000,
			ProductCategoryID: 3,
		},
		{
			Name:              "Powerbank",
			Description:       "Powerbank",
			Price:             100000,
			ProductCategoryID: 4,
		},
		{
			Name:              "Masker",
			Description:       "Masker",
			Price:             5000,
			ProductCategoryID: 5,
		},
		{
			Name:              "Bola Sepak",
			Description:       "Bola Sepak",
			Price:             100000,
			ProductCategoryID: 6,
		},
		{
			Name:              "Lipstik",
			Description:       "Lipstik",
			Price:             50000,
			ProductCategoryID: 7,
		},
		{
			Name:              "Meja",
			Description:       "Meja",
			Price:             500000,
			ProductCategoryID: 8,
		},
		{
			Name:              "Lego",
			Description:       "Lego",
			Price:             100000,
			ProductCategoryID: 9,
		},
		{
			Name:              "Buku",
			Description:       "Buku",
			Price:             50000,
			ProductCategoryID: 10,
		},
	}

	for _, product := range products {
		err := config.DB.Create(&product).Error
		if err != nil {
			log.Printf("Failed to seed Product %s: %v\n", product.Name, err)
		} else {
			log.Printf("Porduct %s seeded successfully\n", product.Name)
		}
	}
}
