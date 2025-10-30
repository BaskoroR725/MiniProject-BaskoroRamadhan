package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"evermos-mini/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// GET /produk
func GetAllProduk(c *fiber.Ctx) error {
	var produk []models.Produk
	config.DB.Preload("Toko").Preload("Category").Find(&produk)
	return c.JSON(fiber.Map{"status": true, "data": produk})
}

// GET /produk/:id
func GetProdukByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var produk models.Produk
	if err := config.DB.Preload("Toko").Preload("Category").First(&produk, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Produk tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"status": true, "data": produk})
}

// GET /produk/search?nama=xxx
func SearchProduk(c *fiber.Ctx) error {
	nama := c.Query("nama")
	var produk []models.Produk
	config.DB.Preload("Toko").Preload("Category").Where("LOWER(nama_produk) LIKE ?", "%"+strings.ToLower(nama)+"%").Find(&produk)
	return c.JSON(fiber.Map{"status": true, "data": produk})
}

// POST /produk
func CreateProduk(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	// cari toko milik user
	var toko models.Toko
	if err := config.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Toko tidak ditemukan"})
	}

	var input struct {
		NamaProduk    string  `json:"nama_produk"`
		HargaReseller float64 `json:"harga_reseller"`
		HargaKonsumen float64 `json:"harga_konsumen"`
		Stok          int     `json:"stok"`
		Deskripsi     string  `json:"deskripsi"`
		CategoryID    uint    `json:"category_id"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	produk := models.Produk{
		NamaProduk:    input.NamaProduk,
		Slug:          strings.ToLower(strings.ReplaceAll(input.NamaProduk, " ", "-")),
		HargaReseller: input.HargaReseller,
		HargaKonsumen: input.HargaKonsumen,
		Stok:          input.Stok,
		Deskripsi:     input.Deskripsi,
		TokoID:        toko.ID,
		CategoryID:    input.CategoryID,
	}

	if err := config.DB.Create(&produk).Error; err != nil {
	return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal menambahkan produk: " + err.Error(),
		})
	}


	//simpan log otomatis
	utils.CreateLogProduk(produk)
	// ambil ulang produk + relasi toko & category
	config.DB.Preload("Toko").Preload("Category").First(&produk, produk.ID)
	return c.JSON(fiber.Map{
			"status":  true,
			"message": "Produk berhasil ditambahkan",
			"data":    produk,
	})
}

// PUT /produk/:id
func UpdateProduk(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var toko models.Toko
	if err := config.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Toko tidak ditemukan"})
	}

	var produk models.Produk
	if err := config.DB.First(&produk, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Produk tidak ditemukan"})
	}

	if produk.TokoID != toko.ID {
		return c.Status(403).JSON(fiber.Map{"status": false, "message": "Tidak punya akses mengubah produk ini"})
	}

	var input struct {
		NamaProduk    string  `json:"nama_produk"`
		HargaReseller float64 `json:"harga_reseller"`
		HargaKonsumen float64 `json:"harga_konsumen"`
		Stok          int     `json:"stok"`
		Deskripsi     string  `json:"deskripsi"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	produk.NamaProduk = input.NamaProduk
	produk.HargaReseller = input.HargaReseller
	produk.HargaKonsumen = input.HargaKonsumen
	produk.Stok = input.Stok
	produk.Deskripsi = input.Deskripsi
	config.DB.Save(&produk)

	utils.CreateLogProduk(produk)


	return c.JSON(fiber.Map{"status": true, "message": "Produk berhasil diperbarui", "data": produk})
}

// DELETE /produk/:id
func DeleteProduk(c *fiber.Ctx) error {
	id := c.Params("id")

	// Cek apakah produk ada dulu
	var produk models.Produk
	if err := config.DB.First(&produk, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Produk tidak ditemukan",
		})
	}

	//  Hapus log_produks yang berelasi dengan produk
	if err := config.DB.Unscoped().Where("produk_id = ?", id).Delete(&models.LogProduk{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal menghapus log produk terkait",
			"error":   err.Error(),
		})
	}

	//  Hapus produk itu sendiri (hard delete)
	if err := config.DB.Unscoped().Delete(&models.Produk{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal menghapus produk",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Produk dan log terkait berhasil dihapus permanen",
	})
}
