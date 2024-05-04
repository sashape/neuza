package main

import (
	"github.com/gofiber/fiber/v2"
	"neuza/backend/routes/articles"
)

func main() {
	app := fiber.New()

	app.Get("/", homepage)

	articles.SetupArticleRoutes(app)

	app.Listen(":8000")
}

func homepage(c *fiber.Ctx) error {
	return c.Send([]byte("NEUZA API"))
}
