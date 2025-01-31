package repository

import (
	"database/sql"

	"user-service/internal/models"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) RegisterUser(user *models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	return r.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email, password FROM users WHERE email=$1"
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
