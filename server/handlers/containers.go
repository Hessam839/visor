package handlers

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
	"visor/server"
	"visor/utils"
)

func HandlerContainerList(c *fiber.Ctx) error {
	conts, contsErr := utils.Visor.ListAllContainer()
	if contsErr != nil {
		log.Fatal(contsErr)
	}
	var items []server.Items
	for _, cont := range *conts {
		items = append(items, server.Items{
			Name:    cont.Names[0],
			Created: humanize.Time(time.Unix(cont.Created, 0)),
			Size:    humanize.Bytes(uint64(cont.SizeRootFs)),
			Status:  cont.Status,
			ID:      cont.ID[:10],
		})
	}
	spew.Dump(items)
	return c.Render("containerlist", server.ViewData{Name: "test", Items: items})
}
