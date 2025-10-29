package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTransaksiRoutes(router fiber.Router) {
	transaksi := router.Group("/transaksi", middleware.JWTProtected)

	transaksi.Post("/", controllers.CreateTransaksi)
	transaksi.Get("/", controllers.GetAllTransaksi)
	transaksi.Get("/:id", controllers.GetTransaksiByID)
	transaksi.Put("/:id/status", controllers.UpdateStatusTransaksi)
}