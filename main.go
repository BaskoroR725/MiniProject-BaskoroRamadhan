package main

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"evermos-mini/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi database
	config.InitDB()

	// Auto-migrate semua tabel
	config.DB.AutoMigrate(
		&models.User{},
		&models.Toko{},
		&models.Alamat{},
		&models.Category{},
		&models.Produk{},
		&models.LogProduk{},
		&models.Transaksi{},
		&models.DetailTransaksi{},
	)

	app := fiber.New()

	// Setup routes (semua dikelola di folder routes/)
	routes.SetupRoutes(app)

	// Jalankan server
	fmt.Println("Server running on http://localhost:8080")
	app.Listen(":8080")
}
