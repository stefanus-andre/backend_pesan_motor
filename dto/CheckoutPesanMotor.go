package dto

type CheckoutPesanMotor struct {
	UserID         uint `json:"user_id"`
	MotorcycleID   uint `json:"motorcycle_id"`
	TanggalPinjam  Date `json:"tanggal_pinjam"`
	TanggalKembali Date `json:"tanggal_kembali"`
}
