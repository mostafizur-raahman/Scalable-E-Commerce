package service

import (
	"user-service/internal/models"
	"user-service/internal/repository"
	"user-service/utils"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return s.repo.RegisterUser(user)
}
