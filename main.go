package main

import (
	"github.com/meherkandukuri/GO_RestApiEventsApp/db"
	"github.com/meherkandukuri/GO_RestApiEventsApp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
