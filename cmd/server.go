package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()

	app.Use("api/v1", cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET,POST,PUT,DELETE,PATCH,OPTIONS"},
		AllowHeaders:     []string{"Origin, Content-Type, Authorization, Accept"},
		AllowCredentials: true,
	}))

	app.Get("/*", static.New("../public"))

	router := app.Group("api/v1")

	router.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Backend app server is running...",
		})
	})

	InitializeRoutes(router)

	app.Listen(":8081")
}
