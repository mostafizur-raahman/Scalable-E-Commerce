package service

import (
	"encoding/json"
	"errors"

	"user-service/internal/models"
	"user-service/internal/repository"
	"user-service/utils"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(user *models.User) (string, error) {
	existingUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if existingUser != nil {
		response, _ := json.Marshal(map[string]string{"error": "email already exists"})
		return string(response), errors.New("email already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword

	if err := s.repo.RegisterUser(user); err != nil {
		return "", err
	}

	response, _ := json.Marshal(map[string]string{"message": "User registered successfully"})
	return string(response), nil
}

func (s *userService) Login(user *models.User) (string, error) {
	existingUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if existingUser == nil {
		response, _ := json.Marshal(map[string]string{"error": "invalid email or password"})
		return string(response), errors.New("invalid email or password")
	}

	// Validate password
	if !utils.CheckPasswordHash(user.Password, existingUser.Password) {
		response, _ := json.Marshal(map[string]string{"error": "invalid email or password"})
		return string(response), errors.New("invalid email or password")
	}

	// Generate token (assuming JWT implementation)
	token, err := utils.GenerateJWT(existingUser.ID, existingUser.Email)
	if err != nil {
		return "", err
	}

	response, _ := json.Marshal(map[string]string{
		"message": "Login successful",
		"token":   token,
	})
	return string(response), nil
}
