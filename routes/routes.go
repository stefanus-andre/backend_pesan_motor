package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_backend/controllers"
	"test_backend/database"
	"test_backend/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware("admin"))
	admin.GET("/dashboard", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Message": "Welcome Admin"})
	})

	// Motorcycle routes
	admin.GET("/get-all-data-motorcycle", controllers.GetAllDataMotorcycles)
	admin.POST("/add-data-motorcycle", controllers.CreateDataMotorcycle)
	admin.GET("/get-detail-data-motorcycle/:id", controllers.GetDetailDataMotorcycles)
	admin.PUT("/update-data-motorcycle/:id", controllers.UpdateMotorcycle)
	admin.DELETE("/delete-data-motorcycle/:id", controllers.DeleteDataMotorcycle)

	// User management routes
	admin.GET("/get-all-users", controllers.GetAllUsers)
	admin.GET("/get-user/:id", controllers.GetDetailUser)
	admin.POST("/add-user", controllers.CreateUser)
	admin.PUT("/update-user/:id", controllers.UpdateUser)
	admin.DELETE("/delete-user/:id", controllers.DeleteUser)

	// Role routes
	admin.GET("/get-all-roles", controllers.GetAllRoles)

	// BUAT OBJEK PesanMotorController
	pesanMotorController := controllers.PesanMotorController{DB: database.DB}

	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware("user"))
	user.GET("/dashboard", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Message": "Welcome User"})
	})
	user.POST("/pesan_motor", pesanMotorController.CheckoutPesanMotor)
	user.GET("/get-all-data-motorcycle-user", controllers.GetAllDataMotorcycles)
}
