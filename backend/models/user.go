package models

import (
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

type UserSession struct {
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
