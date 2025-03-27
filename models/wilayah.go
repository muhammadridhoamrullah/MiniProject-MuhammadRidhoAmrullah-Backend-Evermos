package models

import "gorm.io/gorm"

type Provinsi struct {
	gorm.Model
	IDProvinsi string `json:"id"`
	Name       string `json:"name"`
}

type Kota struct {
	gorm.Model
	IDKota string `json:"id"`
	Name   string `json:"name"`
}

type Kecamatan struct {
	gorm.Model
	IDKecamatan string `json:"id"`
	Name        string `json:"name"`
}

type Kelurahan struct {
	gorm.Model
	IDKelurahan string `json:"id"`
	Name        string `json:"name"`
}
