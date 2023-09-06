package routes

import (

	"github.com/gastrader/gohouse/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupApiV1(app *fiber.App, handlers *handlers.Handler){
	v1 := app.Group("/api/v1")
	SetupUserRoutes(v1, handlers)
}