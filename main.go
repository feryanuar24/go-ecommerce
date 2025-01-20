package main

import (
	"go-ecommerce/config"
	"go-ecommerce/routes"
)

func main() {
	config.ConnectDatabase()

	r := routes.SetupRouter()

	r.Run(":8080")
}
