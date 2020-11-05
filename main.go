package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"log"
	"os"
	"time"
	"visor/dock"
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
	v := dock.NewVisor()
	engine := html.New("./view", ".html")

	app := fiber.New(fiber.Config{Views: engine})
	app.Get("/imagelist", handlers.HandlerGetAllImages)

	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
		Output:     os.Stdout,
	}))

	app.Get("/containerlist", func(c *fiber.Ctx) error {
		conts, contsErr := v.ListAllContainer()
		if contsErr != nil {
			log.Fatal(contsErr)
		}
		var items []Items
		for _, cont := range *conts {
			items = append(items, Items{
				Name:    cont.Names[0],
				Created: humanize.Time(time.Unix(cont.Created, 0)),
				Size:    humanize.Bytes(uint64(cont.SizeRootFs)),
				Status:  cont.Status,
				ID:      cont.ID[:10],
			})
		}
		spew.Dump(items)
		return c.Render("containerlist", ViewData{Name: "test", Items: items})
	})

	app.Get("/stat/:id", handlers.HandlerContainerStats)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	log.Fatal(app.Listen(":8080"))
}
