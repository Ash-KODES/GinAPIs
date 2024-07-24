package main

import (
	"log"
	"socialapp/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	// connecting mongo db
	err := utils.ConnectDB()
	if err != nil {
        log.Fatalf("Error initializing MongoDB: %v", err)
    }

	// it have logger and recovery handler by itself
	server := gin.Default();

	// health check 
	server.GET("/",func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"message" : "server is healthy",
		})
	})

	server.Run();
}
