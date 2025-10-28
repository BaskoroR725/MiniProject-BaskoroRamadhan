package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"evermos-mini/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var input struct {
		Nama     string `json:"nama"`
		Email    string `json:"email"`
		NoTelp   string `json:"no_telp"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Invalid input"})
	}

	if input.Nama == "" || input.Email == "" || input.NoTelp == "" || input.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Semua field wajib diisi"})
	}

	hash, _ := utils.HashPassword(input.Password)

	user := models.User{
		NamaUser:     input.Nama,
		Email:        input.Email,
		NoTelp:       input.NoTelp,
		KataSandi:    hash,
		TanggalLahir: time.Now(),
		JenisKelamin: "",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Email atau no_telp sudah digunakan"})
	}

	// buat toko otomatis
	toko := models.Toko{
    NamaToko: "Toko " + input.Nama,
    UserID:   user.ID,
	}
	config.DB.Create(&toko)

	// ambil ulang user + toko-nya biar tampil di response
	config.DB.Preload("Toko").First(&user, user.ID)

	return c.JSON(fiber.Map{
    "status":  true,
    "message": "Registrasi berhasil",
    "data":    user,
	})

}

func Login(c *fiber.Ctx) error {
	var input struct {
		NoTelp   string `json:"no_telp"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Input tidak valid",
		})
	}

	var user models.User
	if err := config.DB.Where("no_telp = ?", input.NoTelp).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Nomor telepon tidak ditemukan",
		})
	}

	if !utils.CheckPasswordHash(input.Password, user.KataSandi) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Password salah",
		})
	}

	token, _ := utils.GenerateJWT(user.ID)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Login berhasil",
		"data": fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := config.DB.Preload("Toko").First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"status": true, "data": user})
}

func UpdateProfile(c *fiber.Ctx) error {
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
	config.DB.Save(&user)

	return c.JSON(fiber.Map{"status": true, "message": "Profil berhasil diperbarui", "data": user})
}
