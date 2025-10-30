package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"github.com/gofiber/fiber/v2"
)

// CREATE ALAMAT
func CreateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var input models.Alamat

	if input.JudulAlamat == "" {
    input.JudulAlamat = "Alamat Utama"
	}
	
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	input.UserID = userID
	if err := config.DB.Create(&input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal menambahkan alamat"})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Alamat berhasil ditambahkan", "data": input})
}

// GET SEMUA ALAMAT USER
func GetAllAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var alamat []models.Alamat
	config.DB.Where("user_id = ?", userID).Find(&alamat)
	return c.JSON(fiber.Map{"status": true, "data": alamat})
}

// GET ALAMAT BY ID
func GetAlamatByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var alamat models.Alamat
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&alamat).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Alamat tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"status": true, "data": alamat})
}

// UPDATE ALAMAT
func UpdateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var alamat models.Alamat
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&alamat).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Alamat tidak ditemukan"})
	}

	var input models.Alamat
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	config.DB.Model(&alamat).Updates(input)
	return c.JSON(fiber.Map{"status": true, "message": "Alamat berhasil diperbarui", "data": alamat})
}

// DELETE ALAMAT
func DeleteAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var alamat models.Alamat
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&alamat).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Alamat tidak ditemukan"})
	}

	config.DB.Delete(&alamat)
	return c.JSON(fiber.Map{"status": true, "message": "Alamat berhasil dihapus"})
}
