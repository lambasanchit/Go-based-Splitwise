package handlers

import (
	"encoding/json"
	"go-based-splitwise/models"
	storage "go-based-splitwise/storages"
	"net/http"
)

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check for duplicate email
	for _, existingUser := range storage.Users {
		if existingUser.Email == user.Email {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
	}

	newUser := models.NewUser(user.Name, user.Email, user.Password)
	storage.Users[newUser.ID] = newUser

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"userId":  newUser.ID,
	})
}
