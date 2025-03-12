package main

import (
	"github.com/backend-boilerplate-template/routes"
	"github.com/gofiber/fiber/v3"
)

func InitializeRoutes(app fiber.Router) {
	routes.AuthRoutes(app)

	// secured routes
	//app.Use(middlewares.AuthMiddleware)
	routes.AccountRoutes(app)
}
