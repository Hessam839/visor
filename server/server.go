package server

import (
	"github.com/gofiber/fiber/v2"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func NewWebServer() *fiber.App {
	engine := html.New("./server/assets/html", ".html")
	//engine := html.NewFileSystem(rice.MustFindBox("./server/assets/html").HTTPBox(), ".html")
	app := fiber.New(fiber.Config{Views: engine})
	//app.Get("/imagelist", handlers.HandlerGetAllImages)

	//app.Use(handlers.Timer())
	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
		Output:     os.Stdout,
	}))
	return app
}
