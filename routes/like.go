package routes

import (
	"socialapp/controllers"
	"socialapp/middlewares"

	"github.com/gin-gonic/gin"
)

func LikeRoutes(router *gin.RouterGroup) {
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/likes", controllers.CreateLike)
		auth.GET("/likes/:id", controllers.GetLike)
		auth.DELETE("/likes/:id", controllers.DeleteLike)
		auth.GET("/likes", controllers.ListLikes)
	}
}
