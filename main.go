package main

import (
	"github.com/gofiber/fiber/v2"
	"goal-tracking/database"
	"goal-tracking/router"
)

func main() {
	app := fiber.New()

	database.ConnectDb()
	router.SetupRotes(app)

	app.Listen(":3000")
}
