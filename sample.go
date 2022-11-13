package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// start server
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
