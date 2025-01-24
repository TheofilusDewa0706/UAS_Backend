package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(router *gin.Engine) {
	commentRoutes := router.Group("/comments")
	{
		commentRoutes.GET("/", middlewares.AuthMiddleware(1, 2), controllers.GetAllComments)
		commentRoutes.POST("/", middlewares.AuthMiddleware(2), controllers.CreateComment)
		commentRoutes.PUT(":id", middlewares.AuthMiddleware(2), controllers.UpdateComment)
		commentRoutes.DELETE(":id", middlewares.AuthMiddleware(1, 2), controllers.DeleteComment)
	}
}
