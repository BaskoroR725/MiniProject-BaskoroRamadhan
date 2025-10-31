package middleware

import (
	"evermos-mini/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// JWTProtected memverifikasi token JWT sebelum lanjut ke handler
func JWTProtected(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Token tidak ditemukan",
		})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	userID, role, err := utils.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Token tidak valid",
		})
	}

	c.Locals("user_id", userID) 
	c.Locals("role", role)
	return c.Next()
}
