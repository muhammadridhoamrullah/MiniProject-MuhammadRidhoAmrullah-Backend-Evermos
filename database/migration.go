package database

import "mini-project-BE-Evermos/models"

func RunMigration() {
	// Membuat tabel user jika belum ada
	DB.AutoMigrate(&models.User{}, &models.Toko{}, &models.Alamat{}, &models.Provinsi{}, &models.Kota{}, &models.Kecamatan{}, &models.Kelurahan{}, &models.Kategori{}, &models.Produk{}, &models.Transaksi{})
}
