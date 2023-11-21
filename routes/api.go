package routes

import (
	"github.com/PatipatCha/jeab_global_service/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupApiRoutes(app *fiber.App) {

	// *Global
	api := app.Group("/api")
	v1 := api.Group("/v1")
	jguard := v1.Group("/jguard")

	jguard.Get("/config", controllers.GetConfigGlobalHandler)

}
