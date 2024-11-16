package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	host     = os.Getenv("host")
	port     = os.Getenv("port")
	user     = os.Getenv("user")
	password = os.Getenv("pasword")
	dbname   = os.Getenv("dbname")
)

var db *gorm.DB

func GetConnections() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s"+"password=%s dbname=%s sslmode=disabel", host, port, user, password, dbname)

	var err error
	db, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connection database.")
	}
}

func Connection() *gorm.DB {
	return db
}