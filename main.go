package main

import (
	"Financial-Transactions/src/db"
	"Financial-Transactions/src/routes"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	db.GetConnections()
	db.CreateDatabase()
	getRoutes()
	router.Run(":8081")

}

func getRoutes() {
	v1 := router.Group("/v1")
	routes.AddUserRoutes(v1)
}
