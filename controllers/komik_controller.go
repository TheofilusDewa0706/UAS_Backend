package controllers

import (
	"backend/config"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Perbolehkan semua origin (ubah sesuai kebutuhan)
	},
}

// WebSocket clients map
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Komik)

// HandleWebSocket godoc
// @Summary Mengelola koneksi WebSocket
// @Description Menyediakan koneksi WebSocket untuk memperbarui stok komik secara real-time
// @Tags WebSocket
// @Produce application/json
// @Success 200 {string} string "Client terhubung"
// @Router /komik/updates [get]
func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Gagal meng-upgrade ke WebSocket: %v", err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	defer delete(clients, conn)

	// Dengarkan pembaruan stok dari channel broadcast
	for {
		komik := <-broadcast
		for client := range clients {
			err := client.WriteJSON(komik)
			if err != nil {
				log.Printf("Gagal mengirim data ke WebSocket: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// Broadcast stok komik terbaru
func BroadcastStokUpdate(komik models.Komik) {
	broadcast <- komik
}

// GetKomik godoc
// @Summary Menampilkan semua data komik
// @Description Mengambil semua data komik dari database
// @Tags Komik
// @Produce application/json
// @Success 200 {array} models.Komik
// @Router /komik [get]
func GetKomik(c *gin.Context) {
	var komik []models.Komik
	if err := config.DB.Find(&komik).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, komik)
}

// CreateKomik godoc
// @Summary Menambahkan komik baru
// @Description Membuat data komik baru di database
// @Tags Komik
// @Accept application/json
// @Produce application/json
// @Param data body models.Komik true "Data Komik"
// @Success 201 {object} models.Komik
// @Router /komik [post]
func CreateKomik(c *gin.Context) {
	var komik models.Komik
	if err := c.ShouldBindJSON(&komik); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&komik)
	c.JSON(http.StatusCreated, komik)
}

// GetKomikByID godoc
// @Summary Menampilkan detail komik berdasarkan ID
// @Description Mengambil data komik dari database menggunakan ID
// @Tags Komik
// @Produce application/json
// @Param id path int true "ID Komik"
// @Success 200 {object} models.Komik
// @Router /komik/{id} [get]
func GetKomikByID(c *gin.Context) {
	id := c.Param("id")
	var komik models.Komik
	if err := config.DB.First(&komik, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, komik)
}

// UpdateKomik godoc
// @Summary Memperbarui data komik
// @Description Memperbarui data komik berdasarkan ID
// @Tags Komik
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID Komik"
// @Param data body models.Komik true "Data Komik yang Diperbarui"
// @Success 200 {object} models.Komik
// @Router /komik/{id} [put]
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

// DeleteKomik godoc
// @Summary Menghapus data komik
// @Description Menghapus data komik berdasarkan ID
// @Tags Komik
// @Param id path int true "ID Komik"
// @Success 200 {string} string "Data berhasil dihapus"
// @Router /komik/{id} [delete]
func DeleteKomik(c *gin.Context) {
	id := c.Param("id")
	var komik models.Komik
	if err := config.DB.First(&komik, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	config.DB.Delete(&komik)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
