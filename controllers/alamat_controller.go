package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"evermos-mini/utils"
	"github.com/gofiber/fiber/v2"
)

// GET /alamat
func GetAllAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var alamat []models.Alamat
	config.DB.Where("user_id = ?", userID).Find(&alamat)

	return c.JSON(fiber.Map{"status": true, "data": alamat})
}

// GET /alamat/:id
func GetAlamatByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var alamat models.Alamat
	if err := config.DB.First(&alamat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "Alamat tidak ditemukan",
		})
	}

	// pastikan alamat milik user sendiri
	if alamat.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "Tidak punya akses ke alamat ini",
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   alamat,
	})
}

// POST /alamat
func CreateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var input struct {
		JudulAlamat   string `json:"judul_alamat"`
		NamaPenerima  string `json:"nama_penerima"`
		NoTelp        string `json:"no_telp"`
		Provinsi      string `json:"provinsi"`
		Kota          string `json:"kota"`
		Kecamatan     string `json:"kecamatan"`
		Kelurahan     string `json:"kelurahan"`
		DetailAlamat  string `json:"detail_alamat"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	alamat := models.Alamat{
		JudulAlamat:  input.JudulAlamat,
		NamaPenerima: input.NamaPenerima,
		NoTelp:       input.NoTelp,
		Provinsi:     input.Provinsi,
		Kota:         input.Kota,
		Kecamatan:    input.Kecamatan,
		Kelurahan:    input.Kelurahan,
		DetailAlamat: input.DetailAlamat,
		UserID:       userID,
	}

	config.DB.Create(&alamat)
	return c.JSON(fiber.Map{"status": true, "message": "Alamat berhasil ditambahkan", "data": alamat})
}

// PUT /alamat/:id
func UpdateAlamat(c *fiber.Ctx) error {
	id := c.Params("id")

	var alamat models.Alamat
	if err := config.DB.First(&alamat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Alamat tidak ditemukan"})
	}

	if !utils.AuthorizeOwner(c, alamat.UserID) {
		return c.Status(403).JSON(fiber.Map{"status": false, "message": "Tidak punya akses untuk mengubah alamat ini"})
	}

	var input struct {
		JudulAlamat   string `json:"judul_alamat"`
		NamaPenerima  string `json:"nama_penerima"`
		NoTelp        string `json:"no_telp"`
		Provinsi      string `json:"provinsi"`
		Kota          string `json:"kota"`
		Kecamatan     string `json:"kecamatan"`
		Kelurahan     string `json:"kelurahan"`
		DetailAlamat  string `json:"detail_alamat"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	alamat.JudulAlamat = input.JudulAlamat
	alamat.NamaPenerima = input.NamaPenerima
	alamat.NoTelp = input.NoTelp
	alamat.Provinsi = input.Provinsi
	alamat.Kota = input.Kota
	alamat.Kecamatan = input.Kecamatan
	alamat.Kelurahan = input.Kelurahan
	alamat.DetailAlamat = input.DetailAlamat

	config.DB.Save(&alamat)
	return c.JSON(fiber.Map{"status": true, "message": "Alamat berhasil diperbarui", "data": alamat})
}

// DELETE /alamat/:id
func DeleteAlamat(c *fiber.Ctx) error {
	id := c.Params("id")

	var alamat models.Alamat
	if err := config.DB.First(&alamat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Alamat tidak ditemukan"})
	}

	if !utils.AuthorizeOwner(c, alamat.UserID) {
		return c.Status(403).JSON(fiber.Map{"status": false, "message": "Tidak punya akses untuk menghapus alamat ini"})
	}

	config.DB.Delete(&alamat)
	return c.JSON(fiber.Map{"status": true, "message": "Alamat berhasil dihapus"})
}
