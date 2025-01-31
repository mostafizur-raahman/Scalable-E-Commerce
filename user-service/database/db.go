package database

import (
	"database/sql"
	"fmt"
	"log"

	"user-service/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB // Global DB variable

// ConnectDatabase initializes the PostgreSQL connection and returns *sql.DB
func ConnectDatabase(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("❌ Database connection failed:", err)
		return nil, err
	}

	// Ensure database connection is alive
	if err := db.Ping(); err != nil {
		log.Println("❌ Unable to reach database:", err)
		return nil, err
	}

	// Assign to global DB variable
	DB = db

	// Ensure users table exists
	if err := createUsersTable(db); err != nil {
		return nil, err
	}

	fmt.Println("✅ Connected to PostgreSQL successfully")
	return db, nil
}

// createUsersTable ensures the users table exists
func createUsersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
	    id SERIAL PRIMARY KEY,
	    name VARCHAR(255) NOT NULL,
	    email VARCHAR(255) UNIQUE NOT NULL,
	    password TEXT NOT NULL
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Println("❌ Error creating users table:", err)
		return err
	}

	fmt.Println("✅ Users table created successfully")
	return nil
}
