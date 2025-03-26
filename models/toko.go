package models

import "gorm.io/gorm"

type Toko struct {
	gorm.Model
	IDUser   uint   `json:"id_user"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	User     *User  `gorm:"foreignKey:IDUser" json:"user,omitempty"`
}

type UpdateToko struct {
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}
