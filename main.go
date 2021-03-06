package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"visor/server"
	"visor/server/handlers"
)

type Items struct {
	Name    string
	Created string
	Size    string
	Status  string
	ID      string
}
type ViewData struct {
	Name  string
	Items []Items
}

func main() {
	app := server.NewWebServer()
	app.Static("/static", "./server/assets", fiber.Static{
		Compress:  true,
		Browse:    false,
		ByteRange: true,
		MaxAge:    1000000,
	})

	containers := app.Group("/containers")
	containers.Get("/list", handlers.HandlerContainerList)
	containers.Get("/prune", handlers.HandlerContainerPrune)
	containers.Get("/stat/:id", handlers.HandlerContainerStats)
	containers.Get("/start/:id", handlers.HandlerContainerStart)
	containers.Get("/stop/:id", handlers.HandlerContainerStop)
	containers.Get("/pause/:id", handlers.HandlerContainerPause)
	containers.Get("/unpause/:id", handlers.HandlerContainerUnPause)
	containers.Get("/kill/:id", handlers.HandlerContainerKill)

	images := app.Group("/images")
	images.Get("/list", handlers.HandlerGetAllImages)
	images.Get("/prune", handlers.HandlerImagePrune)
	images.Get("/remove/:id", handlers.HandlerImageRemove)
	images.Post("/search", handlers.HandlerImageSearch)
	images.Get("/pull/:iname", handlers.HandlerImagePull)

	app.Get("/utilies", func(c *fiber.Ctx) error {
		return c.Render("utility", nil)
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	log.Fatal(app.Listen(":8080"))
}
