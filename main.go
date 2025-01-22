package main

import (
	"flag"
	"go-ecommerce/config"
	"go-ecommerce/migrations"
	"go-ecommerce/routes"
	"go-ecommerce/seeders"
	"log"
)

func main() {
	seed := flag.Bool("seed", false, "Run seeders")
	flag.Parse()

	config.ConnectDatabase()

	if *seed {
		// Jalankan seeder
		log.Println("Running seeders...")
		seeders.RunAllSeeders()
		log.Println("Seeders completed.")
	} else {
		// Jalankan migrasi
		migrations.Migrate()
		log.Println("Migration completed.")
	}

	r := routes.SetupRouter()

	r.Run(":8080")
}
