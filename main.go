package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renatodaltiba/go-rest-api/database"
	"github.com/renatodaltiba/go-rest-api/services"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", status)

	app.Get("/api/bookmark", services.GetAllBookmarks)
	app.Post("/api/bookmark", services.SaveBookMark)
}

func main() {
	app := fiber.New()
	dbErr := database.InitDatabase()

	if dbErr != nil {
		panic(dbErr)
	}

	setupRoutes(app)
	app.Listen(":9090")
}
