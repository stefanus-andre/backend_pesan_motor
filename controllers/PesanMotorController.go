package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"test_backend/models"
	"time"
)

type PesanMotorController struct {
	DB *gorm.DB
}

// CheckoutPesanMotorRequest defines the request format for booking a motorcycle
type CheckoutPesanMotorRequest struct {
	UserID         uint   `json:"user_id" binding:"required"`
	MotorcycleID   uint   `json:"motorcycle_id" binding:"required"`
	TanggalPinjam  string `json:"tanggal_pinjam" binding:"required"`
	TanggalKembali string `json:"tanggal_kembali" binding:"required"`
}

func (ctrl *PesanMotorController) CheckoutPesanMotor(c *gin.Context) {
	var input CheckoutPesanMotorRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse date strings to time.Time
	tanggalPinjam, err := time.Parse("2006-01-02", input.TanggalPinjam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal pinjam tidak valid. Gunakan format YYYY-MM-DD"})
		return
	}

	tanggalKembali, err := time.Parse("2006-01-02", input.TanggalKembali)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal kembali tidak valid. Gunakan format YYYY-MM-DD"})
		return
	}

	var user models.User
	if err := ctrl.DB.Preload("Role").First(&user, input.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if user.Role.Name != "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Hanya user yang bisa memesan motor"})
		return
	}

	var motor models.Motorcycle
	if err := ctrl.DB.First(&motor, input.MotorcycleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Motorcycle not found"})
		return
	}

	// Calculate duration in days
	durasi := uint32(tanggalKembali.Sub(tanggalPinjam).Hours() / 24)
	if durasi < 1 {
		durasi = 1
	}

	// Calculate total price
	totalHarga := durasi * motor.HargaSewaMotor

	pesan := models.PesanMotor{
		UserID:         user.ID,
		MotorcycleID:   motor.ID,
		TanggalPinjam:  tanggalPinjam,
		TanggalKembali: tanggalKembali,
		TotalHargaSewa: totalHarga,
	}

	if err := ctrl.DB.Create(&pesan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pesanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Pesanan berhasil dibuat",
		"data":        pesan,
		"durasi_sewa": durasi,
		"total_harga": totalHarga,
	})
}
