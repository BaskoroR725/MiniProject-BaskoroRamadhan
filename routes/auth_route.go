package routes

import (
	"evermos-mini/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(r fiber.Router) {
	r.Post("/register", controllers.Register)
	r.Post("/login", controllers.Login)
}
