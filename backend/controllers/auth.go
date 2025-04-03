package controllers

import (
	"backend/config"
	"backend/models"
	"backend/services"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := services.CreateUser(req); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := services.Login(req)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	var user models.User
	err := config.DB.QueryRow("SELECT id, name, first_name, email FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Name, &user.FirstName, &user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
