package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"test_backend/controllers"
	"test_backend/database"
	"test_backend/middleware"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)

	// Admin routes
	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RoleMiddleware("admin"))
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin"})
		})

		// Motorcycle routes
		admin.GET("/motorcycles", controllers.GetAllDataMotorcycles)
		admin.POST("/motorcycles", controllers.CreateDataMotorcycle)
		admin.GET("/motorcycles/:id", controllers.GetDetailDataMotorcycles)
		admin.PUT("/motorcycles/:id", controllers.UpdateMotorcycle)
		admin.DELETE("/motorcycles/:id", controllers.DeleteDataMotorcycle)

		// User management
		admin.GET("/users", controllers.GetAllUsers)
		admin.GET("/users/:id", controllers.GetDetailUser)
		admin.POST("/users", controllers.CreateUser)
		admin.PUT("/users/:id", controllers.UpdateUser)
		admin.DELETE("/users/:id", controllers.DeleteUser)

		// Role management
		admin.GET("/roles", controllers.GetAllRoles)
	}

	// User routes
	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware())
	user.Use(middleware.RoleMiddleware("user"))
	{
		user.GET("/dashboard", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome User"})
		})

		// Motorcycle booking
		pesanMotorController := controllers.PesanMotorController{DB: database.DB}
		user.POST("/bookings", pesanMotorController.CheckoutPesanMotor)
		user.GET("/motorcycles", controllers.GetAllDataMotorcycles)
	}
}
