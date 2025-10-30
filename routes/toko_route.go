package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTokoRoutes(app *fiber.App) {
	toko := app.Group("/toko", middleware.JWTProtected)
	toko.Get("/", controllers.GetTokoByUser)
	toko.Put("/", controllers.UpdateToko)
}
