package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"github.com/gofiber/fiber/v2"
	"time"
	"fmt"
)

type TransaksiInput struct {
	AlamatPengiriman uint `json:"alamat_pengiriman"`
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

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

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

		subtotal := float64(item.Kuantitas) * logProduk.Produk.HargaKonsumen
		totalHarga += subtotal

		// kurangi stok produk
		if logProduk.Produk.Stok < item.Kuantitas {
			return c.Status(400).JSON(fiber.Map{"status": false, "message": "Stok produk tidak mencukupi"})
		}
		logProduk.Produk.Stok -= item.Kuantitas
		config.DB.Save(&logProduk.Produk)

		detailList = append(detailList, models.DetailTransaksi{
			LogProdukID: logProduk.ID,
			Kuantitas:   item.Kuantitas,
			HargaTotal:  subtotal,
		})
	}

	// buat invoice unik
	invoice := fmt.Sprintf("INV-%d-%d", userID, time.Now().Unix())

	transaksi := models.Transaksi{
		UserID:           userID,
		AlamatPengiriman: input.AlamatPengiriman,
		HargaTotal:       totalHarga,
		MetodeBayar:      input.MetodeBayar,
		Invoice:          invoice,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	config.DB.Create(&transaksi)

	for i := range detailList {
		detailList[i].TransaksiID = transaksi.ID
	}
	config.DB.Create(&detailList)

	config.DB.Preload("DetailTransaksi.LogProduk.Produk").First(&transaksi, transaksi.ID)
	return c.JSON(fiber.Map{"status": true, "message": "Transaksi berhasil dibuat", "data": transaksi})
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
