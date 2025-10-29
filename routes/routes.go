package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	// Grouping
	api := app.Group("/api")

	SetupAuthRoutes(api.Group("/auth"))

	SetupUserRoutes(api.Group("/user"))

	SetupProdukRoutes(api.Group("/produk"))

	SetupTransaksiRoutes(api.Group("/transaksi"))
}
