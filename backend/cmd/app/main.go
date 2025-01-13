package main

import (
	"backend/internal/config"
	"backend/internal/interface/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	routes.RegisterRoutes(r)

	r.Run()
}
