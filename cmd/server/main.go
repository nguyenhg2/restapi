package main

import (
	"github.com/gin-gonic/gin"

	"restapi/config"
	"restapi/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
