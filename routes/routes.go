package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	SetupAuthRoutes(app)
	SetupKategoriRoutes(app)
	SetupAlamatRoutes(app)
	SetupTokoRoutes(app)
	SetupProductRoutes(app)
	SetupTrxRoutes(app)
	SetupProvCityRoutes(app)
}
