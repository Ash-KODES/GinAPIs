package main

import (
	
	"socialapp/controllers"
	"socialapp/middlewares"

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

	// Public routes
	server.POST("/register", controllers.Register)
	server.POST("/login", controllers.Login)

	// Protected routes
	protected := server.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/posts", controllers.CreatePost)
		protected.GET("/posts/:id", controllers.GetPost)
		protected.PUT("/posts/:id", controllers.UpdatePost)
		protected.DELETE("/posts/:id", controllers.DeletePost)
		protected.GET("/posts", controllers.ListPosts)

		protected.POST("/comments", controllers.CreateComment)
		protected.GET("/comments/:id", controllers.GetComment)
		protected.PUT("/comments/:id", controllers.UpdateComment)
		protected.DELETE("/comments/:id", controllers.DeleteComment)
		protected.GET("/comments", controllers.ListComments)

		protected.POST("/likes", controllers.CreateLike)
		protected.GET("/likes/:id", controllers.GetLike)
		protected.DELETE("/likes/:id", controllers.DeleteLike)
		protected.GET("/likes", controllers.ListLikes)
	}

	server.Run()
}