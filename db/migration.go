package db

import (
	"fmt"
	"log"
)

func MigrateDb() error {
	migrateQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255) PRIMARY KEY,
  		username VARCHAR(255) NOT NULL UNIQUE,
  		password VARCHAR(255) NOT NULL,
  		email VARCHAR(255) NOT NULL UNIQUE
		refresh_token TEXT
	)
	`

	_, err := DB.Exec(migrateQuery)
	if err != nil {
		log.Println("error on database migration: ",err)
		return err
	}

	fmt.Println("database migration complete")
	return nil
}