package routes

import (
	"github.com/gastrader/gohouse/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(grp fiber.Router, handlers *handlers.Handler){
	useRoute := grp.Group("/user")
	useRoute.Post("/register", handlers.UserRegister)
}