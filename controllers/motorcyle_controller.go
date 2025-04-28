package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	namaMotor := c.PostForm("nama_motor")
	jenisMotor := c.PostForm("jenis_motor")
	nomorPlatMotor := c.PostForm("nomor_plat_motor")
	qtyMotor := c.PostForm("qty_motor")
	hargaSewaMotor := c.PostForm("harga_sewa_motor")
	// tanggalPinjam := c.PostForm("tanggal_pinjam")
	// tanggalKembali := c.PostForm("tanggal_kembali")

	file, err := c.FormFile("image_motor")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	imagePath := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	qtyMotorValue, err := strconv.ParseUint(qtyMotor, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Qty Motor must be a number"})
		return
	}
	hargaSewaMotorValue, err := strconv.ParseUint(hargaSewaMotor, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Harga Sewa Motor must be a number"})
		return
	}

	motorcycle := models.Motorcycle{
		NamaMotor:      namaMotor,
		JenisMotor:     jenisMotor,
		NomorPlatMotor: nomorPlatMotor,
		QtyMotor:       uint32(qtyMotorValue),
		HargaSewaMotor: uint32(hargaSewaMotorValue),
		ImageMotor:     imagePath, // Save the image path
	}

	if err := database.DB.Create(&motorcycle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create motorcycle data"})
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
	QtyMotor       uint32 `json:"qty_motor"` // Changed from string to uint32
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
