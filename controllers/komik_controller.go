package controllers

import (
	"backend/config"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all komik
func GetKomik(c *gin.Context) {
	var komik []models.Komik
	// Coba ambil semua data dari tabel "komik"
	if err := config.DB.Find(&komik).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kirim data ke Postman sebagai JSON
	log.Println("Handler /komik berhasil diakses")
	c.JSON(http.StatusOK, komik)
}

// Create new komik
func CreateKomik(c *gin.Context) {
	var komik models.Komik
	if err := c.ShouldBindJSON(&komik); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&komik)
	c.JSON(http.StatusCreated, komik)
}

// Get komik by ID
func GetKomikByID(c *gin.Context) {
	id := c.Param("id")
	var komik models.Komik
	if err := config.DB.First(&komik, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, komik)
}

// Update komik
func UpdateKomik(c *gin.Context) {
	id := c.Param("id")
	var komik models.Komik
	if err := config.DB.First(&komik, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	if err := c.ShouldBindJSON(&komik); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&komik)
	c.JSON(http.StatusOK, komik)
}

// Delete komik
func DeleteKomik(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Komik{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
