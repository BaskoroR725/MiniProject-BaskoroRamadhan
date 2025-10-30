package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/users", middleware.JWTProtected)
	user.Get("/profile", controllers.GetProfile)
	user.Put("/profile", controllers.UpdateProfile)
}
