package main

import (
	"evermos-mini/config"
	"evermos-mini/models"
	"evermos-mini/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()

	// auto migrate tabel
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
	routes.SetupRoutes(app)

	fmt.Println("🚀 Server running on http://localhost:8080")
	app.Listen(":8080")
}
