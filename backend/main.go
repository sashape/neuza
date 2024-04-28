package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", hello)

	app.Listen(8000)
}

func hello(c *fiber.Ctx) {
	c.Send("Hello world 6")
	// hi
}
