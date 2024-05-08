package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n17ali/events-rest-api/db"
	"github.com/n17ali/events-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.ResigterRoutes(server)

	server.Run(":8080")
}
