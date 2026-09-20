package main

import (
	database "backend/databases"
	"backend/handler"
	"backend/middleware"
	"backend/repositories"
	"backend/services"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal("database connection failed:", err)
	}

	defer db.Close()
	router := gin.Default()
	userRepo := repositories.NewUserRepository(db)

	serviceRepo := repositories.NewServiceRepository(db)

	customerRepo := repositories.NewCustomerRepository(db)

	appointmentRepo := repositories.NewAppointments(db)
	userService := services.NewUserService(
		userRepo,
	)

	serviceService := services.NewServiceService(
		serviceRepo,
		userRepo,
	)

	customerService := services.NewCustomerService(
		customerRepo,
		userRepo,
	)

	appointmentService := services.NewAppointmentService(
		appointmentRepo,
		customerRepo,
		serviceRepo,
	)

	// --------------------------------
	// Handlers
	// --------------------------------

	userHandler := handler.NewUserHanlder(
		userService,
	)

	serviceHandler := handler.NewServiceHandler(
		serviceService,
	)

	customerHandler := handler.NewCustomerHandler(
		customerService,
	)

	appointmentHandler := handler.NewAppointmentHandler(
		appointmentService,
	)

	// --------------------------------
	// Public Routes
	// --------------------------------

	auth := router.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}

	// --------------------------------
	// Protected Routes
	// --------------------------------

	api := router.Group("/api")
	api.Use(middleware.RequireAuth)

	{
		// --------------------------------
		// Profile
		// --------------------------------

		api.GET("/profile", func(c *gin.Context) {

			userID, exists := c.Get("user_id")

			if !exists {
				c.JSON(401, gin.H{
					"error": "user not found",
				})
				return
			}

			c.JSON(200, gin.H{
				"user_id": userID,
			})
		})

		// --------------------------------
		// Services
		// --------------------------------

		api.POST("/services", serviceHandler.Create)

		api.GET("/services", serviceHandler.GetAll)

		api.GET("/services/:id", serviceHandler.GetByID)

		api.PUT("/services/:id", serviceHandler.Update)

		api.DELETE("/services/:id", serviceHandler.Delete)

		// --------------------------------
		// Customers
		// --------------------------------

		api.POST("/customers", customerHandler.Create)

		api.GET("/customers", customerHandler.GetAll)

		api.PUT("/customers/:id", customerHandler.Update)

		api.DELETE("/customers/:id", customerHandler.Delete)

		// --------------------------------
		// Appointments
		// --------------------------------

		api.POST(
			"/appointments",
			appointmentHandler.Create,
		)

		api.GET(
			"/appointments/customer/:customer_id",
			appointmentHandler.GetByCustomer,
		)

		api.GET(
			"/appointments/service/:service_id",
			appointmentHandler.GetByService,
		)

		api.PUT(
			"/appointments/:id/cancel",
			appointmentHandler.Cancel,
		)
	}

	// --------------------------------
	// Server
	// --------------------------------

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
