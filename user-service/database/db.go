package database

import (
	"database/sql"
	"fmt"
	"log"

	"user-service/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Unable to reach database:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully")

	// Automatically create users table if it doesn't exist
	createUsersTable()
}

// createUsersTable ensures the users table exists
func createUsersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
	    id SERIAL PRIMARY KEY,
	    name VARCHAR(255) NOT NULL,
	    email VARCHAR(255) UNIQUE NOT NULL,
	    password TEXT NOT NULL
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("❌ Error creating users table:", err)
	}

	fmt.Println("✅ Users table is ready")
}
