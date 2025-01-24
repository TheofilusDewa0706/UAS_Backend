package main

import (
	"backend/config"
	_ "backend/docs"
	"backend/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Komik API
// @version 1.0
// @description API untuk manajemen komik dan komentar
// @termsOfService http://example.com/terms/
// @contact.name Support Team
// @contact.email support@example.com
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {
	// Membuat instance Gin
	router := gin.Default()

	// Nonaktifkan redirect trailing slash
	router.RedirectTrailingSlash = false

	// Middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                             // URL frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                      // Metode HTTP yang diizinkan
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // Header yang diizinkan
		AllowCredentials: true,                                                          // Jika menggunakan cookie atau header Authorization
		ExposeHeaders:    []string{"Content-Length"},                                    // Header yang dapat diakses oleh client
		MaxAge:           12 * time.Hour,                                                // Cache header selama 12 jam
	}))

	// Koneksi ke database
	setupDatabase()

	// Registrasi routes
	routes.RegisterRoutes(router)
	routes.RegisterCommentRoutes(router) // Aktifkan rute komentar

	// Tambahkan Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Menjalankan server di port 8081
	log.Println("Server berjalan di http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}

// setupDatabase mengatur koneksi ke database
func setupDatabase() {
	log.Println("Menghubungkan ke database...")
	config.ConnectDatabase()
	log.Println("Berhasil terhubung ke database!")
}
