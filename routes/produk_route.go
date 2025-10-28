package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupProdukRoutes(r fiber.Router) {
	// Produk public (tidak butuh login)
	r.Get("/", controllers.GetAllProduk)
	r.Get("/search", controllers.SearchProduk)
	r.Get("/:id", controllers.GetProdukByID)

	// Produk protected (butuh login)
	r.Post("/", middleware.JWTProtected, controllers.CreateProduk)
	r.Put("/:id", middleware.JWTProtected, controllers.UpdateProduk)
	r.Delete("/:id", middleware.JWTProtected, controllers.DeleteProduk)
}
