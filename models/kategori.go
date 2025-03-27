package models

import "gorm.io/gorm"

type Kategori struct {
	gorm.Model
	NamaKategori string   `json:"nama_kategori"`
	Produk       []Produk `gorm:"foreignKey:IDKategori" json:"produk,omitempty"`
}

type InputUpdateKategori struct {
	NamaKategori string `json:"nama_kategori"`
}
