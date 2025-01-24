package websocket

import (
	"log"
	"net/http"
	"sync"

	"backend/config"
	"backend/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan KomikUpdate)
var mu sync.Mutex // Untuk memastikan thread-safe

// Struktur untuk pesan WebSocket
type KomikUpdate struct {
	KomikID uint   `json:"komik_id"`
	Action  string `json:"action"` // "tambah" atau "kurang"
	UserID  uint   `json:"user_id"`
}

// HandleWebSocket untuk menangani koneksi WebSocket
func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Gagal upgrade ke WebSocket:", err)
		return
	}
	defer conn.Close()

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	log.Println("Client terhubung")

	for {
		var update KomikUpdate
		err := conn.ReadJSON(&update)
		if err != nil {
			log.Println("Error membaca pesan:", err)
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			break
		}

		// Kirim data ke channel broadcast
		broadcast <- update
	}
}

// HandleMessages untuk memproses update stok komik
func HandleMessages() {
	for {
		update := <-broadcast

		var komik models.Komik
		if err := config.DB.First(&komik, update.KomikID).Error; err != nil {
			log.Println("Komik tidak ditemukan:", err)
			continue
		}

		mu.Lock()
		if update.Action == "tambah" {
			// Tambahkan stok jika stok tidak melebihi batas awal stok
			komik.Stok++
		} else if update.Action == "kurang" {
			// Kurangi stok jika stok lebih besar dari 0
			if komik.Stok > 0 {
				komik.Stok--
			}
		}
		config.DB.Save(&komik)
		mu.Unlock()

		// Broadcast data stok terbaru ke semua klien
		for client := range clients {
			err := client.WriteJSON(komik)
			if err != nil {
				log.Println("Error mengirim pesan:", err)
				client.Close()
				mu.Lock()
				delete(clients, client)
				mu.Unlock()
			}
		}
	}
}
