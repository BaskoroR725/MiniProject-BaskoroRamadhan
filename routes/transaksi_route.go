package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTransaksiRoutes(router fiber.Router) {
	// semua route transaksi wajib login
	transaksi := router.Group("/", middleware.JWTProtected)

	transaksi.Post("/", controllers.CreateTransaksi)
	transaksi.Get("/", controllers.GetAllTransaksi)
}
