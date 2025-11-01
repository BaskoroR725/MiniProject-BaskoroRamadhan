package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"evermos-mini/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ======================= REGISTER =======================
func Register(c *fiber.Ctx) error {
	var input struct {
		Nama         string `json:"nama" validate:"required,min=3"`
		Email        string `json:"email" validate:"required,email"`
		NoTelp       string `json:"no_telp" validate:"required"`
		KataSandi    string `json:"kata_sandi" validate:"required,min=6"`
		TanggalLahir string `json:"tanggal_lahir"`
		Pekerjaan    string `json:"pekerjaan"`
		IDProvinsi   string `json:"id_provinsi"`
		IDKota       string `json:"id_kota"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	// Validasi otomatis
	if !utils.ValidateStruct(c, input) {
		return nil
	}

	hash, _ := utils.HashPassword(input.KataSandi)

	// Default tanggal lahir jika kosong
	var tglLahir time.Time
	if input.TanggalLahir != "" {
		tgl, err := time.Parse("02/01/2006", input.TanggalLahir)
		if err == nil {
			tglLahir = tgl
		} else {
			tglLahir = time.Now()
		}
	} else {
		tglLahir = time.Now()
	}

	user := models.User{
		NamaUser:     input.Nama,
		Email:        input.Email,
		NoTelp:       input.NoTelp,
		KataSandi:    hash,
		TanggalLahir: tglLahir,
		Role:         "user",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Email atau no_telp sudah digunakan",
		})
	}

	// Buat toko otomatis
	toko := models.Toko{
		NamaToko: "Toko " + input.Nama,
		UserID:   user.ID,
	}
	config.DB.Create(&toko)

	// Simpan alamat dasar dari ID Provinsi dan Kota (jika dikirim)
	if input.IDProvinsi != "" || input.IDKota != "" {
		alamat := models.Alamat{
			NamaPenerima: user.NamaUser,
			NoTelp:       user.NoTelp,
			DetailAlamat: "Alamat utama",
			UserID:       user.ID,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		config.DB.Create(&alamat)
	}

	config.DB.Preload("Toko").Preload("Alamat").First(&user, user.ID)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Registrasi berhasil",
		"data":    user,
	})
}

// ======================= LOGIN =======================
func Login(c *fiber.Ctx) error {
	var input struct {
		NoTelp    string `json:"no_telp" validate:"required"`
		KataSandi string `json:"kata_sandi" validate:"required"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	if !utils.ValidateStruct(c, input) {
		return nil
	}

	var user models.User
	if err := config.DB.Where("no_telp = ?", input.NoTelp).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Nomor telepon tidak ditemukan",
		})
	}

	if !utils.CheckPasswordHash(input.KataSandi, user.KataSandi) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Password salah",
		})
	}

	token, _ := utils.GenerateJWT(user.ID, user.Role)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Login berhasil",
		"data": fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

// ======================= GET PROFILE =======================
func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := config.DB.Preload("Toko").Preload("Alamat").First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"status": true, "data": user})
}

// ======================= UPDATE PROFILE =======================
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

	if input.TanggalLahir != "" {
		tgl, _ := time.Parse("2006-01-02", input.TanggalLahir)
		user.TanggalLahir = tgl
	}

	config.DB.Save(&user)

	return c.JSON(fiber.Map{"status": true, "message": "Profil berhasil diperbarui", "data": user})
}
