package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"test_backend/database"
	"test_backend/handlers"
	"test_backend/middleware"
	"test_backend/models"
	"test_backend/routes"
	"time"
)

func main() {
	database.ConnectDatabase()

	database.DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Motorcycle{}, &models.PesanMotor{})

	var count int64
	database.DB.Model(&models.Role{}).Where("name = ?", "admin").Count(&count)
	if count == 0 {
		log.Println("🔧 Seeding data awal...")

		// Permissions
		viewPerm := models.Permission{Name: "view_data"}
		editPerm := models.Permission{Name: "edit_data"}
		database.DB.FirstOrCreate(&viewPerm, viewPerm)
		database.DB.FirstOrCreate(&editPerm, editPerm)

		adminRole := models.Role{
			Name:        "admin",
			Permissions: []models.Permission{viewPerm, editPerm},
		}
		database.DB.Create(&adminRole)

		userRole := models.Role{
			Name:        "user",
			Permissions: []models.Permission{viewPerm},
		}
		database.DB.Create(&userRole)

		adminPassword, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.MinCost)
		userPassword, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.MinCost)

		adminUser := models.User{
			Email:    "stefanus@example.com",
			Password: string(adminPassword),
			RoleID:   adminRole.ID,
		}
		database.DB.Create(&adminUser)

		// Regular User
		regularUser := models.User{
			Email:    "budi@example.com",
			Password: string(userPassword),
			RoleID:   userRole.ID,
		}
		database.DB.Create(&regularUser)

		log.Println("✅ Seeding selesai.")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/protected", middleware.RBAC("view_data"), handlers.Protected)

	routes.SetupRoutes(r)

	r.Run(":8080")
}
