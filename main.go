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

	app.Get("/containerlist", handlers.HandlerContainerList)

	app.Get("/stat/:id", handlers.HandlerContainerStats)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	log.Fatal(app.Listen(":8080"))
}
