package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "admin:admin123@tcp(database-1.cv6oi4oimtxt.ap-southeast-1.rds.amazonaws.com:3306)/komik?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	log.Println("Berhasil terhubung ke database!")
}
