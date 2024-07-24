package routes

import (
	"socialapp/controllers"
	"socialapp/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.RouterGroup) {
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/comments", controllers.CreateComment)
		auth.GET("/comments/:id", controllers.GetComment)
		auth.PUT("/comments/:id", controllers.UpdateComment)
		auth.DELETE("/comments/:id", controllers.DeleteComment)
		auth.GET("/comments", controllers.ListComments)
	}
}
