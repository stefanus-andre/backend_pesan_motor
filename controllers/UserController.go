package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"test_backend/database"
	"test_backend/models"
)

// GetAllUsers retrieves all users from the database
func GetAllUsers(c *gin.Context) {
	var users []models.User

	result := database.DB.Preload("Role").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch users",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

// GetDetailUser retrieves a specific user by ID
func GetDetailUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := database.DB.Preload("Role").First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User not found",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

// CreateUser adds a new user to the database
func CreateUser(c *gin.Context) {
	var input struct {
		NamaLengkap string `json:"nama_lengkap" binding:"required"`
		NoTelp      string `json:"no_telp" binding:"required"`
		NIK         string `json:"nik" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		BlackList   string `json:"black_list"`
		Password    string `json:"password" binding:"required"`
		RoleID      uint   `json:"roleID" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	// Check if email already exists
	var existingUser models.User
	if database.DB.Where("email = ?", input.Email).First(&existingUser).RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Email already in use",
		})
		return
	}

	// Check if NIK already exists
	if database.DB.Where("nik = ?", input.NIK).First(&existingUser).RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "NIK already registered",
		})
		return
	}

	// Set default blacklist status if not provided
	if input.BlackList == "" {
		input.BlackList = "No"
	}

	// Create new user
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
		return
	}

	user := models.User{
		NamaLengkap: input.NamaLengkap,
		NoTelp:      input.NoTelp,
		NIK:         input.NIK,
		Email:       input.Email,
		BlackList:   input.BlackList,
		Password:    hashedPassword,
		RoleID:      input.RoleID,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create user",
			"error":   result.Error.Error(),
		})
		return
	}

	// Return the created user without the password
	user.Password = ""
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User created successfully",
		"data":    user,
	})
}

// UpdateUser updates a user's information
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Find the user by ID
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User not found",
			"error":   result.Error.Error(),
		})
		return
	}

	// Parse the request body
	var input struct {
		NamaLengkap string `json:"nama_lengkap"`
		NoTelp      string `json:"no_telp"`
		NIK         string `json:"nik"`
		Email       string `json:"email"`
		BlackList   string `json:"black_list"`
		Password    string `json:"password"`
		RoleID      uint   `json:"roleID"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	// Check for email uniqueness if changed
	if input.Email != "" && input.Email != user.Email {
		var existingUser models.User
		if database.DB.Where("email = ?", input.Email).First(&existingUser).RowsAffected > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Email already in use",
			})
			return
		}
		user.Email = input.Email
	}

	// Check for NIK uniqueness if changed
	if input.NIK != "" && input.NIK != user.NIK {
		var existingUser models.User
		if database.DB.Where("nik = ?", input.NIK).First(&existingUser).RowsAffected > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "NIK already registered",
			})
			return
		}
		user.NIK = input.NIK
	}

	// Update fields if provided
	if input.NamaLengkap != "" {
		user.NamaLengkap = input.NamaLengkap
	}
	if input.NoTelp != "" {
		user.NoTelp = input.NoTelp
	}
	if input.BlackList != "" {
		user.BlackList = input.BlackList
	}
	if input.RoleID != 0 {
		user.RoleID = input.RoleID
	}

	// Update password if provided
	if input.Password != "" {
		hashedPassword, err := HashPassword(input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Failed to hash password",
				"error":   err.Error(),
			})
			return
		}
		user.Password = hashedPassword
	}

	// Save the updated user
	result = database.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to update user",
			"error":   result.Error.Error(),
		})
		return
	}

	// Return the updated user without the password
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User updated successfully",
		"data":    user,
	})
}

// DeleteUser removes a user from the database
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid ID format",
		})
		return
	}

	// Delete the user
	result := database.DB.Delete(&models.User{}, idInt)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to delete user",
			"error":   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User deleted successfully",
	})
}

// GetAllRoles retrieves all roles from the database
func GetAllRoles(c *gin.Context) {
	var roles []models.Role

	result := database.DB.Find(&roles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch roles",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   roles,
	})
}

// HashPassword generates a bcrypt hash of the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
