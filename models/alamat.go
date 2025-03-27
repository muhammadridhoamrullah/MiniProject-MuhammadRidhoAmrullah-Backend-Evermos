package models

import "gorm.io/gorm"

type Alamat struct {
	gorm.Model
	IDUser       uint   `json:"id_user"`
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
	IDProvinsi   string `json:"id_provinsi"`
	IDKota       string `json:"id_kota"`
	IDKecamatan  string `json:"id_kecamatan"`
	IDKelurahan  string `json:"id_kelurahan"`

	// Relasi dengan tabel wilayah
	Provinsi  Provinsi  `gorm:"foreignKey:IDProvinsi" json:"provinsi,omitempty"`
	Kota      Kota      `gorm:"foreignKey:IDKota" json:"kota,omitempty"`
	Kecamatan Kecamatan `gorm:"foreignKey:IDKecamatan" json:"kecamatan,omitempty"`
	Kelurahan Kelurahan `gorm:"foreignKey:IDKelurahan" json:"kelurahan,omitempty"`

	User *User `gorm:"foreignKey:IDUser" json:"user,omitempty"`
}

type InputUpdateAlamat struct {
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
	IDProvinsi   string `json:"id_provinsi"`
	IDKota       string `json:"id_kota"`
	IDKecamatan  string `json:"id_kecamatan"`
	IDKelurahan  string `json:"id_kelurahan"`
}
