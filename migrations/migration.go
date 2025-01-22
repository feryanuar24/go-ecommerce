package migrations

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
)

func Migrate() {
	db := config.DB

	// Jalankan migrasi tabel
	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.ProductCategory{},
		&models.Transaction{},
		&models.TransactionItem{},
	)
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}
