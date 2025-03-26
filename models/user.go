package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama         string `json:"nama"`
	KataSandi    string `json:"kata_sandi"`
	NoTelp       string `json:"no_telp" gorm:"unique"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email" gorm:"unique"`
	IDProvinsi   string `json:"id_provinsi"`
	IDKota       string `json:"id_kota"`
	IsAdmin      bool   `json:"is_admin"`
	Toko         Toko   `json:"toko,omitempty" gorm:"foreignKey:IDUser;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type LoginUser struct {
	Email     string `json:"email"`
	KataSandi string `json:"kata_sandi"`
}

type UpdateProfileUser struct {
	Nama         string `json:"nama"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	IDProvinsi   string `json:"id_provinsi"`
	IDKota       string `json:"id_kota"`
}
