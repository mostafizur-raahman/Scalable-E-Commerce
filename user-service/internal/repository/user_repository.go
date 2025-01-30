package repository

import (
	"user-service/database"
	"user-service/internal/models"
)

func RegisterUser(user *models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"

	err := database.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)

	return err
}
