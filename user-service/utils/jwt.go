package utils

import (
	"time"

	"user-service/config"

	"github.com/golang-jwt/jwt/v5"
)

// Function to get JWT Secret only when needed
func getJWTSecret() []byte {
	return []byte(config.LoadConfig().JWTSecret)
}

// Claims struct to include user details
type Claims struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(userID int, email string) (string, error) {
	cfg := config.LoadConfig() // Load configuration once

	// Convert JWTExpiry to time.Duration (default to 24 hours if not set)
	expiryDuration, err := time.ParseDuration(cfg.JWTExpiry)
	if err != nil {
		expiryDuration = 24 * time.Hour // Default to 24 hours
	}

	expirationTime := time.Now().Add(expiryDuration)

	// Create claims with user details
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(getJWTSecret())
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
