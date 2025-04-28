package models

import (
	"gorm.io/gorm"
	"time"
)

type Motorcycle struct {
	gorm.Model
	NamaMotor      string    `json:"nama_motor"`
	JenisMotor     string    `json:"jenis_motor"`
	NomorPlatMotor string    `json:"nomor_plat_motor"`
	QtyMotor       uint32    `json:"qty_motor"`
	HargaSewaMotor uint32    `json:"harga_sewa_motor"`
	TanggalPinjam  time.Time `json:"tanggal_pinjam" gorm:"type:date"`
	TanggalKembali time.Time `json:"tanggal_kembali" gorm:"type:date"`
	ImageMotor     string    `gorm:"size:255" json:"image_motor"` // Add this field
}
