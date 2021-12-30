package main

import (
	"github.com/gin-gonic/gin"
	"passport/router"
)

func main() {
	server := gin.Default()
	router.RegisterRoutes(server)
	server.Run()
}
