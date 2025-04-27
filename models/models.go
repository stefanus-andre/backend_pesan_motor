package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NamaLengkap string `json:"nama_lengkap"`
	NoTelp      string `json:"no_telp"`
	NIK         string `json:"nik"`
	Email       string `json:"email"`
	Image       string `json:"image"`
	BlackList   string `json:"black_list"`
	Password    string
	RoleID      uint
	Role        Role
}
type Role struct {
	gorm.Model
	Name        string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
type Permission struct {
	gorm.Model
	Name string
}
