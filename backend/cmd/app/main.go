package main

import (
	"backend/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	r.Run()
}
