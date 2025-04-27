package models

import (
	"gorm.io/gorm"
	"time"
)

type PesanMotor struct {
	gorm.Model
	UserID         uint       `json:"user_id"`
	User           User       `gorm:"foreignKey:UserID"`
	MotorcycleID   uint       `json:"motorcycle_id"`
	Motorcycle     Motorcycle `gorm:"foreignKey:MotorcycleID"`
	TanggalPinjam  time.Time  `json:"tanggal_pinjam"`
	TanggalKembali time.Time  `json:"tanggal_kembali"`
	TotalHargaSewa uint32     `json:"total_harga_sewa"`
}
