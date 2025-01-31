package service

import "user-service/internal/models"

type UserService interface {
	RegisterUser(user *models.User) error
}
