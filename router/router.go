package router

import (
	"github.com/gofiber/fiber/v2"
	"goal-tracking/handlers"
)

func SetupRotes(app *fiber.App) {

	api := app.Group("/api")

	user := api.Group("/user")
	user.Get("/", handlers.UserGetAll)
	user.Get("/:id", handlers.UserGetById)
	user.Post("/", handlers.UserCreate)
	user.Delete("/:id", handlers.UserDelete)
}
