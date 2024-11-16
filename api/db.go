package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func openDB() (*sql.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set in the environment")
	}
	db, err := sql.Open(
		"postgres",
		databaseURL,
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return db, nil
}

func ConnectToDB() (*sql.DB, error) {

	connection, err := openDB()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to Postgres!")
	return connection, nil
}
