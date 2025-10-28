package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(r fiber.Router) {
	r.Get("/profile", middleware.JWTProtected, controllers.GetProfile)
	r.Put("/profile", middleware.JWTProtected, controllers.UpdateProfile)
}
