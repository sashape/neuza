package main

import (
	"neuza/backend/routes"
	"neuza/backend/database"

	"github.com/gofiber/fiber/v2"
)

const (
	Articles routes.Route = "articles"
)

func main() {
	app := fiber.New()
	database, err := database.Connect()
	if err != nil {
		panic(err)
	}
	app.Get("/", homepage)

	// routes.SetupArticleRoutes(app, database)
	routes.SetupRoutes(app, database, Articles)
	app.Listen(":8000")
}

func homepage(c *fiber.Ctx) error {
	return c.Send([]byte("NEUZA API"))
}
