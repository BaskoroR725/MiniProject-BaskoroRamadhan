package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupProdukRoutes(app *fiber.App) {
	app.Get("/produk", controllers.GetAllProduk)
	app.Get("/produk/:id", controllers.GetProdukByID)

	produk := app.Group("/produk", middleware.JWTProtected)
	produk.Post("/", controllers.CreateProduk)
	produk.Put("/:id", controllers.UpdateProduk)
	produk.Delete("/:id", controllers.DeleteProduk)
	produk.Post("/upload", controllers.UploadGambarProduk)

}
