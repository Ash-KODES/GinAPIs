package main

import (
	"socialapp/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	
	// It has logger and recovery handler by itself
	server := gin.Default()

	// Health check
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is healthy",
		})
	})

	routes.SetupRouter(server)
	server.Run()
}