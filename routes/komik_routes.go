package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Rute login
	r.POST("/login", controllers.Login)
	// Rute CRUD komik
	komik := r.Group("/komik")
	{
		komik.GET("/", middlewares.AuthMiddleware(1, 2), controllers.GetKomik)
		komik.POST("/", middlewares.AuthMiddleware(1), controllers.CreateKomik)
		komik.GET("/:id", middlewares.AuthMiddleware(1, 2), controllers.GetKomikByID)
		komik.PUT("/:id", middlewares.AuthMiddleware(1), controllers.UpdateKomik)
		komik.DELETE("/:id", middlewares.AuthMiddleware(1), controllers.DeleteKomik)
	}
}
