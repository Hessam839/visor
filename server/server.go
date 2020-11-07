package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"os"
	"visor/server/handlers"
)

func NewWebServer() *fiber.App {
	engine := html.New("./server/assets/html", ".html")

	app := fiber.New(fiber.Config{Views: engine})
	app.Get("/imagelist", handlers.HandlerGetAllImages)

	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
		Output:     os.Stdout,
	}))
	return app
}
