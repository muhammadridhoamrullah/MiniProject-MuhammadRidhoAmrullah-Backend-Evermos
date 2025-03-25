package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Variabel global untuk akses ke DB
var DB *gorm.DB

// Fungsi untuk menghubungkan ke database
func ConnectDB() {

	// DSN (Data Source Name) untuk menghubungkan ke database
	dsn := "root:@tcp(127.0.0.1:3306)/mini_project?charset=utf8mb4&parseTime=True&loc=Local"

	// Membuat koneksi ke database menggunakan GORM
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Jika terjadi error saat menghubungkan ke database
	if err != nil {
		log.Fatal(("Gagal terhubung ke database"))
	}

	// Mengatur variabel global DB dengan database yang telah terhubung
	DB = database
	log.Println("Berhasil terhubung ke database")
}