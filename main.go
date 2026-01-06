package main

import (
	"log"
	"mini-indobat-backend/config"
	"mini-indobat-backend/database"
	"mini-indobat-backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.LoadConfig()
	database.ConnectDatabase(cfg)

	router := routes.SetupRoutes()
	router.Run(":8080")
}
