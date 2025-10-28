package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Auth
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)

	// User
	user := app.Group("/user", middleware.JWTProtected)
	user.Get("/profile", controllers.GetProfile)
	user.Put("/profile", controllers.UpdateProfile)

	// Produk public
	app.Get("/produk", controllers.GetAllProduk)
	app.Get("/produk/:id", controllers.GetProdukByID)
	app.Get("/produk/search", controllers.SearchProduk)

	// Produk protected
	produk := app.Group("/produk", middleware.JWTProtected)
	produk.Post("/", controllers.CreateProduk)
	produk.Put("/:id", controllers.UpdateProduk)
	produk.Delete("/:id", controllers.DeleteProduk)
}
