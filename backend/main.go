package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Vibhuair20/Coupon-System-MVP/config"
	"github.com/Vibhuair20/Coupon-System-MVP/handlers"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/coupons/applicable", handlers.GetApplicableCoupons)
	app.Post("/coupons/validate", handlers.ValidateCoupon)
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Initialize database
	config.InitDB()

	// Create Fiber app
	app := fiber.New()

	// Setup routes
	SetupRoutes(app)

	// Get port from environment variable
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
