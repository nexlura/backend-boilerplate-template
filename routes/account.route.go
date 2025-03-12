package routes

import (
	"github.com/backend-boilerplate-template/controllers"
	"github.com/gofiber/fiber/v3"
)

func AccountRoutes(app fiber.Router) {

	app.Post("/accounts", func(c fiber.Ctx) error {
		return controllers.CreateAccount(c)
	})

	//app.Get("/accounts", func(c *fiber.Ctx) error {
	//	return controllers.GetAllAccounts(c)
	//})
	//
	//app.Patch("/accounts/:id", func(c *fiber.Ctx) error {
	//	return controllers.UpdateAccount(c)
	//})
	//
	//app.Get("/accounts/:id/properties", func(c *fiber.Ctx) error {
	//	return controllers.GetPropertyByUserId(c)
	//})
	//
	//app.Get("/accounts/:id", func(c *fiber.Ctx) error {
	//	return controllers.GetAccount(c)
	//})
	//
	//app.Delete("/accounts/:id", func(c *fiber.Ctx) error {
	//	return controllers.DeleteAccount(c)
	//})
}
