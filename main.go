package main

import (
	"log"
	"socialapp/controllers"
	// "socialapp/middlewares"
	"socialapp/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// it have logger and recovery handler by itself
	server := gin.Default()

	// health check
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is healthy",
		})
	})

	// connecting mongo db
	err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Error initializing MongoDB: %v", err)
	}

	// Public routes
	server.POST("/register", controllers.Register)
	// server.POST("/login", controllers.Login)

	// server.Use(middlewares.AuthMiddleware())
	// {
	// 	server.POST("/posts", controllers.CreatePost)
	// 	server.GET("/posts/:id", controllers.GetPost)
	// 	server.PUT("/posts/:id", controllers.UpdatePost)
	// 	server.DELETE("/posts/:id", controllers.DeletePost)
	// 	server.GET("/posts", controllers.ListPosts)

	// 	server.POST("/comments", controllers.CreateComment)
	// 	server.GET("/comments/:id", controllers.GetComment)
	// 	server.PUT("/comments/:id", controllers.UpdateComment)
	// 	server.DELETE("/comments/:id", controllers.DeleteComment)
	// 	server.GET("/comments", controllers.ListComments)

	// 	server.POST("/likes", controllers.CreateLike)
	// 	server.GET("/likes/:id", controllers.GetLike)
	// 	server.DELETE("/likes/:id", controllers.DeleteLike)
	// 	server.GET("/likes", controllers.ListLikes)
	// }

	server.Run()
}
