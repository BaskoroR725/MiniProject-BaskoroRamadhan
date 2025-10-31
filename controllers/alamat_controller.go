package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"github.com/gofiber/fiber/v2"
)

// GET semua alamat user
func GetAllAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var alamat []models.Alamat
	if err := config.DB.Where("user_id = ?", userID).Find(&alamat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal mengambil alamat",
		})
	}

	return c.JSON(fiber.Map{"status": true, "data": alamat})
}

// GET alamat by ID
func GetAlamatByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var alamat models.Alamat
	if err := config.DB.First(&alamat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Alamat tidak ditemukan"})
	}

	if alamat.UserID != userID {
		return c.Status(403).JSON(fiber.Map{"status": false, "message": "Tidak punya akses ke alamat ini"})
	}

	return c.JSON(fiber.Map{"status": true, "data": alamat})
}

// POST tambah alamat baru
func CreateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var input models.Alamat
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	input.UserID = userID

	if err := config.DB.Create(&input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal menambah alamat"})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Alamat berhasil ditambahkan", "data": input})
}

// PUT update alamat
func UpdateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var alamat models.Alamat
	if err := config.DB.First(&alamat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Alamat tidak ditemukan"})
	}

	if alamat.UserID != userID {
		return c.Status(403).JSON(fiber.Map{"status": false, "message": "Tidak punya akses ke alamat ini"})
	}

	var input models.Alamat
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	config.DB.Model(&alamat).Updates(input)
	return c.JSON(fiber.Map{"status": true, "message": "Alamat berhasil diperbarui", "data": alamat})
}

// DELETE hapus alamat
func DeleteAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var alamat models.Alamat
	if err := config.DB.First(&alamat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Alamat tidak ditemukan"})
	}

	if alamat.UserID != userID {
		return c.Status(403).JSON(fiber.Map{"status": false, "message": "Tidak punya akses ke alamat ini"})
	}

	config.DB.Delete(&alamat)
	return c.JSON(fiber.Map{"status": true, "message": "Alamat berhasil dihapus"})
}
