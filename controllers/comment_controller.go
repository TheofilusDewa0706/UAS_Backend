package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllComments godoc
// @Summary Menampilkan semua komentar
// @Description Admin dapat melihat semua komentar, sedangkan user hanya dapat melihat komentarnya sendiri
// @Tags Komentar
// @Produce application/json
// @Success 200 {array} models.Comment
// @Router /comments [get]
// @Security BearerAuth
func GetAllComments(c *gin.Context) {
	role, _ := c.Get("role_id")

	if role == 1 {
		// Admin dapat melihat semua komentar
		var comments []models.Comment
		if err := config.DB.Find(&comments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, comments)
		return
	}

	if role == 2 {
		// User hanya dapat melihat komentar miliknya sendiri
		userID, _ := c.Get("user_id")
		var comments []models.Comment
		if err := config.DB.Where("user_id = ?", userID).Find(&comments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, comments)
		return
	}

	c.JSON(http.StatusForbidden, gin.H{"error": "Role tidak diizinkan untuk melihat komentar"})
}

// CreateComment godoc
// @Summary Membuat komentar baru
// @Description User dapat membuat komentar baru
// @Tags Komentar
// @Accept application/json
// @Produce application/json
// @Param data body models.Comment true "Data Komentar"
// @Success 201 {object} models.Comment
// @Router /comments [post]
// @Security BearerAuth
func CreateComment(c *gin.Context) {
	role, _ := c.Get("role_id")
	if role != 2 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Hanya user yang dapat membuat komentar"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil user_id dari context
	userID, _ := c.Get("user_id")
	comment.UserID = userID.(uint)

	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

// UpdateComment godoc
// @Summary Memperbarui komentar
// @Description User dapat memperbarui komentarnya sendiri
// @Tags Komentar
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID Komentar"
// @Param data body models.Comment true "Data Komentar yang Diperbarui"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [put]
// @Security BearerAuth
func UpdateComment(c *gin.Context) {
	role, _ := c.Get("role_id")
	if role != 2 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Hanya user yang dapat mengedit komentar"})
		return
	}

	id := c.Param("id")
	var comment models.Comment

	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Komentar tidak ditemukan"})
		return
	}

	userID, _ := c.Get("user_id")
	if comment.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Tidak diizinkan mengedit komentar ini"})
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&comment)
	c.JSON(http.StatusOK, comment)
}

// DeleteComment godoc
// @Summary Menghapus komentar
// @Description Admin dapat menghapus komentar siapa saja, sedangkan user hanya dapat menghapus komentarnya sendiri
// @Tags Komentar
// @Param id path int true "ID Komentar"
// @Success 200 {string} string "Komentar berhasil dihapus"
// @Router /comments/{id} [delete]
// @Security BearerAuth
func DeleteComment(c *gin.Context) {
	role, _ := c.Get("role_id")
	id := c.Param("id")

	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Komentar tidak ditemukan"})
		return
	}

	userID, _ := c.Get("user_id")

	if role == 2 {
		// User hanya bisa menghapus komentarnya sendiri
		if comment.UserID != userID.(uint) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Tidak diizinkan menghapus komentar ini"})
			return
		}
	}

	// Admin dapat menghapus komentar siapa saja
	config.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{"message": "Komentar berhasil dihapus"})
}
