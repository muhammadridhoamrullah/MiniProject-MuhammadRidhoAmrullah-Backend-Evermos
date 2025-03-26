package models

import "gorm.io/gorm"

type Alamat struct {
	gorm.Model
	IDUser       uint   `json:"id_user"`
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}
