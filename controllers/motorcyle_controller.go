package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_backend/database"
	"test_backend/models"
	"time"
)

type CreateMotorcycleRequest struct {
	NamaMotor      string `json:"nama_motor"`
	JenisMotor     string `json:"jenis_motor"`
	NomorPlatMotor string `json:"nomor_plat_motor"`
	QtyMotor       uint32 `json:"qty_motor"`
	HargaSewaMotor uint32 `json:"harga_sewa_motor"`
	TanggalPinjam  string `json:"tanggal_pinjam"`
	TanggalKembali string `json:"tanggal_kembali"`
	ImageMotor     string `json:"image_motor"`
	UserID         uint   `json:"user_id"`
}

func CreateDataMotorcycle(c *gin.Context) {
	var request CreateMotorcycleRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//tanggalPinjam, err := time.Parse("2006-01-02", request.TanggalPinjam)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal pinjam tidak valid. Gunakan format YYYY-MM-DD"})
	//	return
	//}
	//
	//tanggalKembali, err := time.Parse("2006-01-02", request.TanggalKembali)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal kembali tidak valid. Gunakan format YYYY-MM-DD"})
	//	return
	//}

	motorcycle := models.Motorcycle{
		NamaMotor:      request.NamaMotor,
		JenisMotor:     request.JenisMotor,
		NomorPlatMotor: request.NomorPlatMotor,
		QtyMotor:       request.QtyMotor,
		HargaSewaMotor: request.HargaSewaMotor,
		ImageMotor:     request.ImageMotor,
	}

	if err := database.DB.Create(&motorcycle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Membuat Data Motor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"motorcycle": motorcycle})
}

func GetAllDataMotorcycles(c *gin.Context) {
	var motorcycles []models.Motorcycle

	if err := database.DB.Find(&motorcycles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Melihat data motorcycles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data semua motorcycles": motorcycles})
}

func GetDetailDataMotorcycles(c *gin.Context) {
	id := c.Param("id")
	var motorcycle models.Motorcycle

	if err := database.DB.First(&motorcycle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gagal Membuat Data Motorcycles"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"motorcycle": motorcycle})
}

type UpdateMotorcycleRequest struct {
	NamaMotor      string `json:"nama_motor"`
	JenisMotor     string `json:"jenis_motor"`
	NomorPlatMotor string `json:"nomor_plat_motor"`
	QtyMotor       string `json:"qty_motor"`
	HargaSewaMotor uint32 `json:"harga_sewa_motor"`
	TanggalPinjam  string `json:"tanggal_pinjam"`
	TanggalKembali string `json:"tanggal_kembali"`
	ImageMotor     string `json:"image_motor"`
}

func UpdateMotorcycle(c *gin.Context) {
	id := c.Param("id")
	var motorcycle models.Motorcycle

	if err := database.DB.First(&motorcycle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data Motorcycles tidak ditemukan"})
		return
	}

	var request UpdateMotorcycleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields that are provided
	if request.NamaMotor != "" {
		motorcycle.NamaMotor = request.NamaMotor
	}

	if request.JenisMotor != "" {
		motorcycle.JenisMotor = request.JenisMotor
	}

	if request.NomorPlatMotor != "" {
		motorcycle.NomorPlatMotor = request.NomorPlatMotor
	}

	//if request.QtyMotor != "" {
	//	motorcycle.QtyMotor = request.QtyMotor
	//}

	if request.HargaSewaMotor != 0 {
		motorcycle.HargaSewaMotor = request.HargaSewaMotor
	}

	if request.TanggalPinjam != "" {
		tanggalPinjam, err := time.Parse("2006-01-02", request.TanggalPinjam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal pinjam tidak valid. Gunakan format YYYY-MM-DD"})
			return
		}
		motorcycle.TanggalPinjam = tanggalPinjam
	}

	if request.TanggalKembali != "" {
		tanggalKembali, err := time.Parse("2006-01-02", request.TanggalKembali)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal kembali tidak valid. Gunakan format YYYY-MM-DD"})
			return
		}
		motorcycle.TanggalKembali = tanggalKembali
	}

	if request.ImageMotor != "" {
		motorcycle.ImageMotor = request.ImageMotor
	}

	if err := database.DB.Save(&motorcycle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data motorcycle"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data motorcycle sudah update": motorcycle})
}

func DeleteDataMotorcycle(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var motorcycle models.Motorcycle
	if err := database.DB.First(&motorcycle, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Delete(&motorcycle)

	c.JSON(http.StatusOK, gin.H{"message": "Data motorcycle sudah dihapus"})
}
