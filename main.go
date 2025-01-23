package main

import (
	"backend/config"
	"backend/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	// Menjalankan server di port 8081
	log.Println("Server berjalan di http://localhost:8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}

// setupDatabase mengatur koneksi ke database
func setupDatabase() {
	log.Println("Menghubungkan ke database...")
	config.ConnectDatabase()
	log.Println("Berhasil terhubung ke database!")
}
