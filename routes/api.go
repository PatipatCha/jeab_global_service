package routes

import (
	"github.com/PatipatCha/jeab_global_service/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupApiRoutes(app *fiber.App, store *session.Store) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Azure Functions!")
	})

	// * VMS
	// vms := app.Group("/api")
	// vms.Get("/create-vms-table", controllers.CreateVmsTable)
	// vms.Post("/create-vms", controllers.CreateVms)
	// vms.Get("/get-vms", controllers.GetVms)

	// *Global
	globalGroup := app.Group("/api")
	globalGroup.Get("/get-config-jeabguard", controllers.GetConfigGlobal)

}
