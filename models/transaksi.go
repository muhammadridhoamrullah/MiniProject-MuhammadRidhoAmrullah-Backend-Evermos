package models

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	IDUser           uint   `json:"id_user"`
	AlamatPengiriman string `json:"alamat_pengiriman"`
	HargaTotal       int    `json:"harga_total"`
	KodeInvoice      string `json:"kode_invoice"`
	MethodBayar      string `json:"method_bayar"`
	Status           string `json:"status"`

	// Relasi dengan tabel user
	User User `gorm:"foreignKey:IDUser" json:"user,omitempty"`
}

type InputEditTransaksi struct {
	AlamatPengiriman string `json:"alamat_pengiriman"`
	HargaTotal       int    `json:"harga_total"`
	MethodBayar      string `json:"method_bayar"`
	Status           string `json:"status"`
}
