package database

import (
	"api/src/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConectionDB)
	if err != nil {
		log.Fatalf("[database] [msg: Error to connect %v]", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Printf("[database] [msg: Error to ping %v]", err)
	}

	return db, nil
}
