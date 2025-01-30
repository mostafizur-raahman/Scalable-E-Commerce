package service

import (
	"user-service/internal/models"
	"user-service/internal/repository"
	"user-service/utils"
)

func RegisterUser(user *models.User) error {
	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPass

	return repository.RegisterUser(user)
}
