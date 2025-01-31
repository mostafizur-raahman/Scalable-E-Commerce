package main

import (
	"fmt"
	"log"
	"net/http"

	"user-service/config"
	"user-service/database"
	"user-service/internal/controller"
	"user-service/internal/repository"
	"user-service/internal/routes"
	"user-service/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	// Try loading `.env` manually
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: No .env file found, using system environment variables")
	}

	cfg := config.LoadConfig()

	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal("❌ Failed to initialize database:", err)
	}
	defer db.Close()

	// Initialize repository, service, and controller
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Create a new ServeMux
	mux := http.NewServeMux()
	routes.UserRoutes(mux, userController)

	// Start the HTTP server
	serverAddr := ":" + cfg.ServerPort
	fmt.Printf("✅ User service is running on port %s\n", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(serverAddr, mux))
}
