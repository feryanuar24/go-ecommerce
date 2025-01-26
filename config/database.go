package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection based on the current environment.
func ConnectDatabase() {
	// Pastikan working directory ke root proyek
	os.Chdir("..") // Sesuaikan jika `go test` dijalankan dari subdirektori

	// Tentukan environment file (default ke .env)
	envFile := ".env"
	if os.Getenv("APP_ENV") == "testing" {
		envFile = ".env.testing"
	}

	// Load environment variables
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	// Baca konfigurasi database dari environment
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Hubungkan ke database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Assign instance ke variabel global
	DB = database
}
