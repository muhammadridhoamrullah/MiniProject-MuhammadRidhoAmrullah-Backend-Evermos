package models

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	NamaProduk    string `json:"nama_produk"`
	Slug          string `json:"slug"`
	HargaReseller string `json:"harga_reseller"`
	HargaKonsumen string `json:"harga_konsumen"`
	Stok          int    `json:"stok"`
	Deskripsi     string `json:"deskripsi"`
	IDToko        uint   `json:"id_toko"`
	IDKategori    uint   `json:"id_kategori"`

	// Relasi dengan tabel toko dan kategori
	Toko     Toko     `gorm:"foreignKey:IDToko" json:"toko,omitempty"`
	Kategori Kategori `gorm:"foreignKey:IDKategori" json:"kategori,omitempty"`
}
