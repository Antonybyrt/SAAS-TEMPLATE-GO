package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(localhost:8889)/SAAS")
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	return nil
}
