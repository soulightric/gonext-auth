package middleware

import (
	"strings"

	"fiber-auth/utils"

	"github.com/gofiber/fiber/v2"
)

// Protected adalah middleware untuk memvalidasi JWT pada route yang memerlukan autentikasi.
// Cara pakai: taruh sebelum handler, contoh: app.Get("/me", middleware.Protected(), handlers.Me)
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token tidak ditemukan, silakan login",
			})
		}

		// Header biasanya berformat: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Format token tidak valid",
			})
		}

		claims, err := utils.ParseJWT(parts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token tidak valid atau sudah kedaluwarsa",
			})
		}

		// user_id dari JWT claims (float64 karena JSON numbers di-parse sebagai float64)
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token tidak valid",
			})
		}

		// Simpan user_id ke context, bisa diakses di handler berikutnya
		c.Locals("user_id", uint(userID))

		return c.Next()
	}
}
