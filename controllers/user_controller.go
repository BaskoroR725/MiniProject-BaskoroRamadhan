package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GET /user
func GetMyProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := config.DB.Preload("Toko").Preload("Alamat").First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "User tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   user,
	})
}

// PUT /user
func UpdateMyProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var input struct {
		Nama          string `json:"nama"`
		NoTelp        string `json:"no_telp"`
		TanggalLahir  string `json:"tanggal_Lahir"`
		JenisKelamin  string `json:"jenis_kelamin"`
		Pekerjaan     string `json:"pekerjaan"`
		Email         string `json:"email"`
		IDProvinsi    string `json:"id_provinsi"`
		IDKota        string `json:"id_kota"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Input tidak valid",
		})
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "User tidak ditemukan",
		})
	}

	// Update field sesuai input
	user.NamaUser = input.Nama
	user.NoTelp = input.NoTelp
	user.JenisKelamin = input.JenisKelamin
	user.Pekerjaan = input.Pekerjaan
	user.Email = input.Email
	user.IDProvinsi = input.IDProvinsi
	user.IDKota = input.IDKota

	// Parse tanggal lahir dari format Postman Rakamin
	if input.TanggalLahir != "" {
		if parsed, err := time.Parse("02/01/2006", input.TanggalLahir); err == nil {
			user.TanggalLahir = parsed
		}
	}

	config.DB.Save(&user)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Profil berhasil diperbarui",
		"data":    user,
	})
}
