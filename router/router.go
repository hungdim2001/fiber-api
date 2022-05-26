package router

import (
	noteRoutes "hungdim2001/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)
func SetupRoutes (app *fiber.App){
	api :=app.Group("/api", logger.New())
	noteRoutes.SetupNoteRoutes(api)
}