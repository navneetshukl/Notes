package main

import (
	"log"
	"notes/internal/adapter/persistence/db"
	"notes/pkg/logging"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
		return
	}
	log := logging.NewLogger()
	dbConnection := db.NewDBConnection(log)
	_,err=dbConnection.ConnectToDB()
}
