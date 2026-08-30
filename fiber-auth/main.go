package main

import (
	"log"
	"os"

	"fiber-auth/database"
	"fiber-auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: file .env tidak ditemukan, menggunakan environment variable sistem")
	}

	// Koneksi & migrasi database
	database.Connect()

	app := fiber.New(fiber.Config{
		AppName: "Fiber Auth API",
	})

	// Middleware global
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // ganti sesuai domain frontend saat production
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Daftarkan semua route
	routes.Setup(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
