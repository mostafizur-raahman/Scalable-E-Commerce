package main

import (
	"fmt"
	"log"
	"net/http"

	"user-service/config"
	"user-service/database"
	"user-service/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Try loading `.env` manually
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: No .env file found, using system environment variables")
	}

	cfg := config.LoadConfig()

	database.ConnectDatabase(cfg)

	routes.UserRoutes()

	serverAddr := ":" + cfg.ServerPort
	fmt.Printf("✅ User service is running on port %s\n", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
