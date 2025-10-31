package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"evermos-mini/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TransaksiInput struct {
	AlamatPengiriman uint   `json:"alamat_pengiriman"`
	MetodeBayar      string `json:"metode_bayar"`
	Items []struct {
		LogProdukID uint `json:"log_produk_id"`
		Kuantitas   int  `json:"kuantitas"`
	} `json:"items"`
}

// POST /api/transaksi
func CreateTransaksi(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var input TransaksiInput

	// parsing input dari body JSON
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	// validasi minimal satu produk
	if len(input.Items) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Tidak ada produk dalam transaksi"})
	}

	var totalHarga float64
	var detailList []models.DetailTransaksi

	for _, item := range input.Items {
		var logProduk models.LogProduk
		if err := config.DB.Preload("Produk").First(&logProduk, item.LogProdukID).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": false, "message": fmt.Sprintf("LogProduk ID %d tidak ditemukan", item.LogProdukID)})
		}

		produk := logProduk.Produk // ambil produk dari log

		// validasi stok
		if produk.Stok < item.Kuantitas {
			return c.Status(400).JSON(fiber.Map{"status": false, "message": "Stok produk tidak mencukupi"})
		}

		// hitung subtotal
		subtotal := float64(item.Kuantitas) * produk.HargaKonsumen
		totalHarga += subtotal

		// update stok dan buat log baru
		produk.Stok -= item.Kuantitas
		config.DB.Save(&produk)
		utils.CreateLogProduk(produk)

		// simpan detail transaksi (sementara belum commit)
		detailList = append(detailList, models.DetailTransaksi{
			LogProdukID: logProduk.ID,
			Kuantitas:   item.Kuantitas,
			HargaTotal:  subtotal,
		})
	}

	
	invoice := fmt.Sprintf("INV-%d-%d", userID, time.Now().Unix())

	transaksi := models.Transaksi{
		UserID:           userID,
		AlamatPengiriman: input.AlamatPengiriman,
		HargaTotal:       totalHarga,
		MetodeBayar:      input.MetodeBayar,
		Invoice:          invoice,
		Status:           "Menunggu Pembayaran",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	// simpan transaksi utama
	if err := config.DB.Create(&transaksi).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": "Gagal menyimpan transaksi"})
	}

	// simpan detail transaksi dengan relasi transaksi ID
	for i := range detailList {
		detailList[i].TransaksiID = transaksi.ID
	}
	config.DB.Create(&detailList)

	// preload relasi untuk response
	config.DB.
		Preload("User.Toko").
		Preload("DetailTransaksi.LogProduk.Produk").
		Preload("DetailTransaksi.LogProduk.Produk.Toko").
		Preload("DetailTransaksi.LogProduk.Produk.Category").
		Preload("DetailTransaksi.LogProduk.Toko").
		Preload("DetailTransaksi.LogProduk.Category").
		First(&transaksi, transaksi.ID)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Transaksi berhasil dibuat",
		"data":    transaksi,
	})
}

// GET /api/transaksi
func GetAllTransaksi(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var transaksi []models.Transaksi
	config.DB.Preload("DetailTransaksi.LogProduk.Produk").
		Where("user_id = ?", userID).
		Find(&transaksi)

	return c.JSON(fiber.Map{"status": true, "data": transaksi})
}

// GET /api/transaksi/:id
func GetTransaksiByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var transaksi models.Transaksi
	if err := config.DB.
		Preload("User.Toko").
		Preload("DetailTransaksi.LogProduk.Produk.Toko").
		Preload("DetailTransaksi.LogProduk.Produk.Category").
		Preload("DetailTransaksi.LogProduk.Toko").
		Preload("DetailTransaksi.LogProduk.Category").
		Where("user_id = ?", userID).
		First(&transaksi, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "Transaksi tidak ditemukan",
		})
	}

	if transaksi.UserID != userID {
    return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
        "status": false,
        "message": "Tidak punya akses ke transaksi ini",
    })
}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   transaksi,
	})
}

// PUT /api/transaksi/:id/status
func UpdateStatusTransaksi(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var input struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	// Validasi status yang diizinkan
	validStatuses := map[string]bool{
		"Menunggu Pembayaran": true,
		"Dikemas":             true,
		"Dikirim":             true,
		"Selesai":             true,
	}
	if !validStatuses[input.Status] {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Status tidak valid",
		})
	}

	var transaksi models.Transaksi
	if err := config.DB.Where("user_id = ?", userID).First(&transaksi, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Transaksi tidak ditemukan"})
	}

	transaksi.Status = input.Status
	transaksi.UpdatedAt = time.Now()
	config.DB.Save(&transaksi)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Status transaksi berhasil diperbarui",
		"data":    transaksi,
	})
}

