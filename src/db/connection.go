package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)


var db *gorm.DB

func GetConnections() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file.")
	}
	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("PORT")
		user     = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname   = os.Getenv("DBNAME")
	)
	
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connection database.")
	}
}

func Connection() *gorm.DB {
	return db
}