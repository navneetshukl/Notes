package db

import (
	"fmt"
	"log"
	"notes/pkg/logging"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	log *logging.Log
}

func NewDBConnection(log *logging.Log) *DBConnection {
	return &DBConnection{log: log}
}
func (db *DBConnection) ConnectToDB() (*gorm.DB, error) {
	DB_USER := os.Getenv("POSTGRES_USER")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_NAME := os.Getenv("POSTGRES_DB")
	log.Println(DB_USER, DB_PASSWORD, DB_NAME)

	dsn := fmt.Sprintf("host=localhost port=5434 user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		db.log.Error("Failed to connect to database", err)
		return nil, err
	}

	db.log.Info("Connected to database")
	return dbConn, nil
}
