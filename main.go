package main

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
	"go-ecommerce/routes"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.ProductCategory{},
		&models.Transaction{},
		&models.TransactionItem{},
	)

	r := routes.SetupRouter()

	r.Run(":8080")
}
