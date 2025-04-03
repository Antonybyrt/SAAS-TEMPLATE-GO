package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	var err error
	dbStringConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	DB, err = sql.Open("mysql", dbStringConnection)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	_, err = DB.Exec("CREATE DATABASE IF NOT EXISTS " + os.Getenv("DB_NAME"))
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	dbStringConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	DB, err = sql.Open("mysql", dbStringConnection)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	return nil
}

func createDB() error {
	_, err := DB.Exec("CREATE DATABASE IF NOT EXISTS SAAS")
	if err != nil {
		return err
	}

	_, err = DB.Exec("USE SAAS")
	if err != nil {
		return err
	}

	return nil
}

func createUserTable() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			first_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			has_upgraded BOOLEAN NOT NULL DEFAULT FALSE
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func createSession() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS user_sessions (
			user_id INT NOT NULL,
			expires_at DATETIME NOT NULL,
			token VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func CreateDatabase() error {
	createDB()
	createUserTable()
	createSession()
	return nil
}
