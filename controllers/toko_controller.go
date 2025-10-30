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
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Toko tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"status": true, "data": toko})
}

// PUT /toko
func UpdateToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var toko models.Toko
	if err := config.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Toko tidak ditemukan"})
	}

	var input struct {
		NamaToko string `json:"nama_toko"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	toko.NamaToko = input.NamaToko
	config.DB.Save(&toko)

	return c.JSON(fiber.Map{"status": true, "message": "Toko berhasil diperbarui", "data": toko})
}
