package routes

import (
	"socialapp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// it have logger and recovery handler by itself
	router := gin.Default()

	// Public routes
	AuthRoutes(router)

	// Protected routes
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		PostRoutes(auth)
		CommentRoutes(auth)
		LikeRoutes(auth)
	}

	return router
}
