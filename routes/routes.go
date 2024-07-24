package routes

import (
	"socialapp/middlewares"
	"socialapp/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(server *gin.Engine) *gin.Engine {
	// Public routes
	server.POST("/register", controllers.Register)
    server.POST("/login", controllers.Login)

    auth := server.Group("/")
    auth.Use(middlewares.AuthMiddleware())
    {
        auth.POST("/posts", controllers.CreatePost)
        auth.GET("/posts/:id", controllers.GetPost)
        auth.PUT("/posts/:id", controllers.UpdatePost)
        auth.DELETE("/posts/:id", controllers.DeletePost)
        auth.GET("/posts", controllers.ListPosts)

        auth.POST("/comments", controllers.CreateComment)
        auth.GET("/comments/:id", controllers.GetComment)
        auth.PUT("/comments/:id", controllers.UpdateComment)
        auth.DELETE("/comments/:id", controllers.DeleteComment)
        auth.GET("/comments", controllers.ListComments)

        auth.POST("/likes", controllers.CreateLike)
        auth.GET("/likes/:id", controllers.GetLike)
        auth.DELETE("/likes/:id", controllers.DeleteLike)
        auth.GET("/likes", controllers.ListLikes)
    }
	return server
}
