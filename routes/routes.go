package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	SetupAuthRoutes(app)
	SetupUserRoutes(app)
	SetupTokoRoutes(app)
	SetupAlamatRoutes(app)
	SetupKategoriRoutes(app)
	SetupProdukRoutes(app)
	SetupTransaksiRoutes(app)
	SetupAlamatRoutes(app)
}
