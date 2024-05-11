package main

import (
	"neuza/backend/routes/articles"
	"neuza/backend/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database, err := database.Connect()
	if err != nil {
		panic(err)
	}
	app.Get("/", homepage)

	routes.SetupArticleRoutes(app, database)

	app.Listen(":8000")
}

func homepage(c *fiber.Ctx) error {
	return c.Send([]byte("NEUZA API"))
}
