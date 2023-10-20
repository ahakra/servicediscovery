package routes

import (
	"github.com/ahakra/servicediscovery/serviceHealth/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func FiberRoutes(app *fiber.App) {

	app.Get("/", controller.GetAllServicesData)
}
