package routes

import (
	"fiber-auth/handlers"
	"fiber-auth/middleware"

	"github.com/gofiber/fiber/v2"
)

// Setup mendaftarkan semua route aplikasi.
func Setup(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/auth")

	// Route publik (tidak perlu login)
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Route yang dilindungi (butuh token JWT)
	auth.Get("/me", middleware.Protected(), handlers.Me)
}
