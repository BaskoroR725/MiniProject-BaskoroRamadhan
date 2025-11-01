package routes

import (
	"evermos-mini/controllers"
	"evermos-mini/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTokoRoutes(app *fiber.App) {
	toko := app.Group("/toko", middleware.JWTProtected)

	toko.Get("/", controllers.GetAllToko)           
	toko.Get("/my", controllers.GetMyToko)          
	toko.Get("/:id_toko", controllers.GetTokoByID)  
	toko.Put("/:id_toko", controllers.UpdateToko)   
}
