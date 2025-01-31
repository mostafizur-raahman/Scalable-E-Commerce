package config

import (
	"os"
)

// Config struct to hold environment variables
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	ServerPort string
	JWTSecret  string
	JWTExpiry  string
}

// LoadConfig loads environment variables
func LoadConfig() *Config {
	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		ServerPort: os.Getenv("SERVER_PORT"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		JWTExpiry:  os.Getenv("JWT_EXPIRY"),
	}

	return cfg
}
