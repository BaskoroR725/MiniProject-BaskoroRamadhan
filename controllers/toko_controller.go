package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"

	"github.com/gofiber/fiber/v2"
)

// GET /toko
func GetTokoByUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var toko models.Toko
	if err := config.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Toko tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   toko,
	})
}

// PUT /toko
func UpdateToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var toko models.Toko
	if err := config.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Toko tidak ditemukan",
		})
	}

	var input struct {
		NamaToko string `json:"nama_toko"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Input tidak valid",
		})
	}

	if input.NamaToko == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Nama toko wajib diisi",
		})
	}

	toko.NamaToko = input.NamaToko
	if err := config.DB.Save(&toko).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal memperbarui toko",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Toko berhasil diperbarui",
		"data":    toko,
	})
}
