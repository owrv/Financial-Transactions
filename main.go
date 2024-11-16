package main

import (
	"Financial-Transactions/src/db"
	"Financial-Transactions/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router = gin.Default()

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file.")
	}
	db.GetConnections()
	getRoutes()
	router.Run(":8081")

}

func getRoutes() {
	v1 := router.Group("/v1")
	routes.AddUserRoutes(v1)
}
