package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllComments - Admin dan User dapat melihat komentar
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

// CreateComment - Hanya user yang dapat membuat komentar
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

// UpdateComment - Hanya user yang dapat mengedit komentar miliknya sendiri
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

// DeleteComment - Admin dapat menghapus komentar siapa saja, User hanya dapat menghapus komentarnya sendiri
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
