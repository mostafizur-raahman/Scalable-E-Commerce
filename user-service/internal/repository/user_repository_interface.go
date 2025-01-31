package repository

import "user-service/internal/models"

type UserRepository interface {
	RegisterUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}
