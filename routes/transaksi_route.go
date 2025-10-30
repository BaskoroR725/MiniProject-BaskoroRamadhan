package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTransaksiRoutes(app *fiber.App) {
	transaksi := app.Group("/transaksi", middleware.JWTProtected)
	transaksi.Get("/", controllers.GetAllTransaksi)
	transaksi.Get("/:id", controllers.GetTransaksiByID)
	transaksi.Post("/", controllers.CreateTransaksi)
	transaksi.Put("/:id", controllers.UpdateStatusTransaksi)
}
