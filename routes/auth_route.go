package routes

import (
	"evermos-mini/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.Register)
}
