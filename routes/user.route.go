package routes

import (
	"github.com/backend-boilerplate-template/controllers"
	"github.com/gofiber/fiber/v3"
)

func UserRoutes(app fiber.Router) {

	app.Post("/users", func(c fiber.Ctx) error {
		return controllers.CreateUser(c)
	})

	app.Get("/users", func(c fiber.Ctx) error {
		return controllers.GetUsers(c)
	})

	app.Get("/users/:param", func(c fiber.Ctx) error {
		return controllers.GetUser(c)
	})

	app.Patch("/users/:id", func(c fiber.Ctx) error {
		return controllers.UpdateUser(c)
	})

	app.Delete("/users/:id", func(c fiber.Ctx) error {
		return controllers.DeleteUser(c)
	})
}
