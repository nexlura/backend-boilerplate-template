package routes

import (
	"github.com/backend-boilerplate-template/controllers"
	"github.com/gofiber/fiber/v3"
)

func AuthRoutes(app fiber.Router) {
	app.Post("/auth/register", func(c fiber.Ctx) error {
		return controllers.CreateUser(c)
	})

	//app.Post("/auth/login", func(c fiber.Ctx) error {
	//	return controllers.Login(c)
	//})
	//
	//app.Get("/auth/logout", func(c fiber.Ctx) error {
	//	return controllers.Logout(c)
	//})
	//
	//app.Post("/auth/forgot-password", func(c fiber.Ctx) error {
	//	return controllers.ForgotPassword(c)
	//})
	//
	//app.Patch("/auth/reset-password", func(c fiber.Ctx) error {
	//	return controllers.ResetPassword(c)
	//})
}
