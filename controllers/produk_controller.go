package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"evermos-mini/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GET /product?nama_produk=&limit=&page=&category_id=&toko_id=&min_harga=&max_harga=
func GetAllProduct(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	nama := c.Query("nama_produk", "")
	categoryID := c.Query("category_id", "")
	tokoID := c.Query("toko_id", "")
	minHarga := c.Query("min_harga", "")
	maxHarga := c.Query("max_harga", "")

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var produk []models.Produk
	query := config.DB.Preload("Toko").Preload("Category")

	if nama != "" {
		query = query.Where("LOWER(nama_produk) LIKE ?", "%"+strings.ToLower(nama)+"%")
	}
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if tokoID != "" {
		query = query.Where("toko_id = ?", tokoID)
	}
	if minHarga != "" {
		query = query.Where("harga_konsumen >= ?", minHarga)
	}
	if maxHarga != "" {
		query = query.Where("harga_konsumen <= ?", maxHarga)
	}

	var total int64
	query.Model(&models.Produk{}).Count(&total)

	if err := query.Offset(offset).Limit(limit).Find(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal mengambil produk"})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return c.JSON(fiber.Map{
		"status": true,
		"pagination": fiber.Map{
			"page":        page,
			"limit":       limit,
			"total_data":  total,
			"total_pages": totalPages,
		},
		"data": produk,
	})
}

// GET /product/:id
func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var produk models.Produk
	if err := config.DB.Preload("Toko").Preload("Category").First(&produk, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Produk tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"status": true, "data": produk})
}

// POST /product
func CreateProduct(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var toko models.Toko
	if err := config.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Toko tidak ditemukan"})
	}

	namaProduk := c.FormValue("nama_produk")
	hargaReseller, _ := strconv.ParseFloat(c.FormValue("harga_reseller"), 64)
	hargaKonsumen, _ := strconv.ParseFloat(c.FormValue("harga_konsumen"), 64)
	stok, _ := strconv.Atoi(c.FormValue("stok"))
	deskripsi := c.FormValue("deskripsi")
	categoryID, _ := strconv.Atoi(c.FormValue("category_id"))

	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Gambar produk wajib diisi"})
	}

	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	savePath := fmt.Sprintf("./uploads/%s", filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal menyimpan gambar"})
	}

	gambarURL := fmt.Sprintf("http://localhost:8080/uploads/%s", filename)

	produk := models.Produk{
		NamaProduk:    namaProduk,
		Slug:          strings.ToLower(strings.ReplaceAll(namaProduk, " ", "-")),
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stok,
		Deskripsi:     deskripsi,
		TokoID:        toko.ID,
		CategoryID:    uint(categoryID),
		Gambar:        gambarURL,
	}

	if err := config.DB.Create(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal menambahkan produk", "error": err.Error()})
	}

	utils.CreateLogProduk(produk)

	return c.JSON(fiber.Map{"status": true, "message": "Produk berhasil ditambahkan", "data": produk})
}

// PUT /product/:id
func UpdateProduct(c *fiber.Ctx) error {
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
		return c.Status(403).JSON(fiber.Map{"status": false, "message": "Tidak punya akses"})
	}

	namaProduk := c.FormValue("nama_produk")
	hargaReseller, _ := strconv.ParseFloat(c.FormValue("harga_reseller"), 64)
	hargaKonsumen, _ := strconv.ParseFloat(c.FormValue("harga_konsumen"), 64)
	stok, _ := strconv.Atoi(c.FormValue("stok"))
	deskripsi := c.FormValue("deskripsi")

	if file, err := c.FormFile("photo"); err == nil {
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		savePath := fmt.Sprintf("./uploads/%s", filename)
		c.SaveFile(file, savePath)
		produk.Gambar = fmt.Sprintf("http://localhost:8080/uploads/%s", filename)
	}

	produk.NamaProduk = namaProduk
	produk.HargaReseller = hargaReseller
	produk.HargaKonsumen = hargaKonsumen
	produk.Stok = stok
	produk.Deskripsi = deskripsi

	config.DB.Save(&produk)
	utils.CreateLogProduk(produk)

	return c.JSON(fiber.Map{"status": true, "message": "Produk berhasil diperbarui", "data": produk})
}

// DELETE /product/:id
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var produk models.Produk
	if err := config.DB.First(&produk, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Produk tidak ditemukan"})
	}

	config.DB.Unscoped().Where("produk_id = ?", id).Delete(&models.LogProduk{})
	config.DB.Unscoped().Delete(&produk)

	return c.JSON(fiber.Map{"status": true, "message": "Produk berhasil dihapus"})
}
