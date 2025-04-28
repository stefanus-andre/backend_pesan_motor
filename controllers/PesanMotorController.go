package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"test_backend/models"
)

type PesanMotorController struct {
	DB *gorm.DB
}

type CheckoutPesanMotorRequest struct {
	MotorcycleID   uint   `json:"motorcycle_id" binding:"required"`
	TanggalPinjam  string `json:"tanggal_pinjam" binding:"required"`
	TanggalKembali string `json:"tanggal_kembali" binding:"required"`
}

func (ctrl *PesanMotorController) CheckoutPesanMotor(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is missing. Please login again."})
		return
	}

	// Get user role
	role, exists := c.Get("role")
	if !exists || role != "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only regular users can book motorcycles"})
		return
	}

	var input CheckoutPesanMotorRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse dates
	tanggalPinjam, err := time.Parse("2006-01-02", input.TanggalPinjam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for tanggal_pinjam. Use YYYY-MM-DD"})
		return
	}

	tanggalKembali, err := time.Parse("2006-01-02", input.TanggalKembali)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for tanggal_kembali. Use YYYY-MM-DD"})
		return
	}

	// Check motorcycle exists
	var motor models.Motorcycle
	if err := ctrl.DB.First(&motor, input.MotorcycleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Motorcycle not found"})
		return
	}

	// Calculate duration
	durasi := uint32(tanggalKembali.Sub(tanggalPinjam).Hours() / 24)
	if durasi < 1 {
		durasi = 1
	}

	// Calculate total price
	totalHarga := durasi * motor.HargaSewaMotor

	// Create booking
	pesan := models.PesanMotor{
		UserID:         userID.(uint),
		MotorcycleID:   motor.ID,
		TanggalPinjam:  tanggalPinjam,
		TanggalKembali: tanggalKembali,
		TotalHargaSewa: totalHarga,
	}

	if err := ctrl.DB.Create(&pesan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Booking created successfully",
		"data":        pesan,
		"durasi_sewa": durasi,
		"total_harga": totalHarga,
	})
}
