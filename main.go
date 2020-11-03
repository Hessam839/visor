package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
	"time"
	"visor/dock"
)

type Items struct {
	Name    string
	Created string
	Size    string
	Status  string
}
type ViewData struct {
	Name  string
	Items []Items
}

func main() {
	v := dock.NewVisor()
	engine := html.New("./view", ".html")

	app := fiber.New(fiber.Config{Views: engine})
	app.Get("/imagelist", func(c *fiber.Ctx) error {
		imgs, imgsErr := v.GetAllImages()
		if imgsErr != nil {
			log.Fatal(imgsErr)
		}
		var images []Items
		for _, img := range *imgs {
			images = append(images, Items{
				Name:    img.RepoTags[0],
				Created: humanize.Time(time.Unix(img.Created, 0)),
				Size:    humanize.Bytes(uint64(img.VirtualSize)),
			})
		}
		spew.Dump(images)
		return c.Render("imagelist", ViewData{Name: "test", Items: images})
	})

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
			})
		}
		spew.Dump(items)
		return c.Render("containerlist", ViewData{Name: "test", Items: items})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	log.Fatal(app.Listen(":8080"))
}
