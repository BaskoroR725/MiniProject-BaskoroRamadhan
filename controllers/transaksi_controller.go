package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"evermos-mini/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// POST /trx
func CreateTrx(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var input struct {
		AlamatPengiriman uint   `json:"alamat_pengiriman"`
		MetodeBayar      string `json:"metode_bayar"`
		Items []struct {
			LogProdukID uint `json:"log_produk_id"`
			Kuantitas   int  `json:"kuantitas"`
		} `json:"items"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	if len(input.Items) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Produk tidak boleh kosong"})
	}

	var totalHarga float64
	var details []models.DetailTransaksi

	for _, item := range input.Items {
		var logProduk models.LogProduk
		if err := config.DB.Preload("Produk").First(&logProduk, item.LogProdukID).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"status": false, "message": fmt.Sprintf("LogProduk %d tidak ditemukan", item.LogProdukID)})
		}

		produk := logProduk.Produk
		if produk.Stok < item.Kuantitas {
			return c.Status(400).JSON(fiber.Map{"status": false, "message": "Stok produk tidak mencukupi"})
		}

		produk.Stok -= item.Kuantitas
		config.DB.Save(&produk)
		utils.CreateLogProduk(produk)

		subtotal := produk.HargaKonsumen * float64(item.Kuantitas)
		totalHarga += subtotal

		details = append(details, models.DetailTransaksi{
			LogProdukID: logProduk.ID,
			Kuantitas:   item.Kuantitas,
			HargaTotal:  subtotal,
		})
	}

	invoice := fmt.Sprintf("INV-%d-%d", userID, time.Now().Unix())

	transaksi := models.Transaksi{
		UserID:           userID,
		AlamatPengiriman: input.AlamatPengiriman,
		MetodeBayar:      input.MetodeBayar,
		HargaTotal:       totalHarga,
		Invoice:          invoice,
		Status:           "Menunggu Pembayaran",
	}

	config.DB.Create(&transaksi)
	for i := range details {
		details[i].TransaksiID = transaksi.ID
	}
	config.DB.Create(&details)

	config.DB.Preload("DetailTransaksi.LogProduk.Produk").First(&transaksi, transaksi.ID)

	return c.JSON(fiber.Map{"status": true, "message": "Transaksi berhasil dibuat", "data": transaksi})
}

// GET /trx
func GetAllTrx(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var transaksi []models.Transaksi

	config.DB.Where("user_id = ?", userID).
		Preload("DetailTransaksi.LogProduk.Produk").
		Find(&transaksi)

	return c.JSON(fiber.Map{"status": true, "data": transaksi})
}

// GET /trx/:id
func GetTrxByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var transaksi models.Transaksi
	if err := config.DB.
		Preload("DetailTransaksi.LogProduk.Produk").
		Where("id = ? AND user_id = ?", id, userID).
		First(&transaksi).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Transaksi tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"status": true, "data": transaksi})
}

// PUT /trx/:id/status
func UpdateStatusTrx(c *fiber.Ctx) error {
	id := c.Params("id")
	var input struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	var transaksi models.Transaksi
	if err := config.DB.First(&transaksi, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Transaksi tidak ditemukan"})
	}

	transaksi.Status = input.Status
	transaksi.UpdatedAt = time.Now()
	config.DB.Save(&transaksi)

	return c.JSON(fiber.Map{"status": true, "message": "Status berhasil diperbarui", "data": transaksi})
}
