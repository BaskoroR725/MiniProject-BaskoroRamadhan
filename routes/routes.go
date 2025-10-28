package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)

	user := app.Group("/user", middleware.JWTProtected)
	user.Get("/profile", controllers.GetProfile)
	user.Put("/profile", controllers.UpdateProfile)
}
