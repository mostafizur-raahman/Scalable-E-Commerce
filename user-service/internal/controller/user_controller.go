package controller

import (
	"encoding/json"
	"net/http"

	"user-service/internal/models"
	"user-service/internal/service"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := service.RegisterUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	// Send user data in the response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user":    userResponse,
	})
}
