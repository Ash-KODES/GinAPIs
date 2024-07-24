package routes

import (
	"socialapp/controllers"
	"socialapp/middlewares"

	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.RouterGroup) {
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/posts", controllers.CreatePost)
		auth.GET("/posts/:id", controllers.GetPost)
		auth.PUT("/posts/:id", controllers.UpdatePost)
		auth.DELETE("/posts/:id", controllers.DeletePost)
		auth.GET("/posts", controllers.ListPosts)
	}
}
