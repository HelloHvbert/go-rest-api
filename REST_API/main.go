package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	db.InitDB()
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)
	
	router.Run(":8080")
}

