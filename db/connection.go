package db

import (
	"database/sql"
	"os"
)

func OpenConnection() (*sql.DB, error) {

	databaseURL := os.Getenv("DATABASE_URL")
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
