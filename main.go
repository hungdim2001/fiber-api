package main

import (
	"hungdim2001/database"
	"hungdim2001/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
    // Start a new fiber app
    app := fiber.New()
    database.ConnectDB()
    // Send a string back for GET calls to the endpoint "/"
    app.Get("/", func(c *fiber.Ctx) error {
        err := c.SendString("And the API is here!")
        return err
    })
    router.SetupRoutes(app)
    // Listen on PORT 3000
    app.Listen(":3000")
}