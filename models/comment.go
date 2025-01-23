package models

type Comment struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserID   uint   `json:"user_id"`  // Relasi ke User
	KomikID  uint   `json:"komik_id"` // Relasi ke Komik
	Komentar string `json:"komentar"`
}
