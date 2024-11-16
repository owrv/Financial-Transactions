package db

import (
	"os"
)

var (
	host     = os.Getenv("host")
	port     = os.Getenv("port")
	user     = os.Getenv("user")
	password = os.Getenv("pasword")
	dbname   = os.Getenv("dbname")
)

var db *gorm.DB

func Connection() *gorm.DB {
	return db
}