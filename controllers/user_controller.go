package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// =================== GET MY PROFILE ===================
func GetMyProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := config.DB.Preload("Toko").Preload("Alamat").First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"status": true, "data": user})
}

// =================== UPDATE PROFILE ===================
func UpdateMyProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var input struct {
		NamaUser     string `json:"nama_user"`
		JenisKelamin string `json:"jenis_kelamin"`
		TanggalLahir string `json:"tanggal_lahir"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User tidak ditemukan"})
	}

	user.NamaUser = input.NamaUser
	user.JenisKelamin = input.JenisKelamin

	if input.TanggalLahir != "" {
		tgl, _ := time.Parse("2006-01-02", input.TanggalLahir)
		user.TanggalLahir = tgl
	}

	config.DB.Save(&user)
	return c.JSON(fiber.Map{"status": true, "message": "Profil berhasil diperbarui", "data": user})
}
