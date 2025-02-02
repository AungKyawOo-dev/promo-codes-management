package main

import (
	"promo_codes_management/config"
	"promo_codes_management/routes"
)

func main() {
	config.ConnectDatabase()
	router := routes.SetupRoutes()
	router.Run(":8080")
}
