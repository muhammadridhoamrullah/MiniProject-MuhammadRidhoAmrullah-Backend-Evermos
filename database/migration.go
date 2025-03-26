package database

import "mini-project-BE-Evermos/models"

func RunMigration() {
	// Membuat tabel user jika belum ada
	DB.AutoMigrate(&models.User{}, &models.Toko{})
}