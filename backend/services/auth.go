package services

import (
	"backend/config"
	"backend/models"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func CreateUser(user models.RegisterRequest) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("INSERT INTO users (name, first_name, email, password) VALUES (?, ?, ?, ?)",
		user.Name, user.FirstName, user.Email, hashedPassword)
	return err
}

func Login(loginReq models.LoginRequest) (string, error) {
	var user models.User
	err := config.DB.QueryRow("SELECT id, email, password FROM users WHERE email = ?", loginReq.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !CheckPasswordHash(loginReq.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := GenerateToken()
	if err != nil {
		return "", err
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	_, err = config.DB.Exec("INSERT INTO user_sessions (user_id, token, expires_at) VALUES (?, ?, ?)",
		user.ID, token, expiresAt)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string) (int, error) {
	var userID int
	var expiresAtStr string
	err := config.DB.QueryRow("SELECT user_id, DATE_FORMAT(expires_at, '%Y-%m-%d %H:%i:%s') FROM user_sessions WHERE token = ?", token).Scan(&userID, &expiresAtStr)
	if err != nil {
		return 0, errors.New("invalid token")
	}

	expiresAt, err := time.Parse("2006-01-02 15:04:05", expiresAtStr)
	if err != nil {
		return 0, errors.New("invalid date format")
	}

	if time.Now().After(expiresAt) {
		return 0, errors.New("token expired")
	}

	return userID, nil
}
