package handlers

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
	"visor/utils"
)

func HandlerContainerList(c *fiber.Ctx) error {
	conts, contsErr := utils.Visor.ContainerListAll()
	if contsErr != nil {
		log.Fatal(contsErr)
		return c.SendString(contsErr.Error())
	}
	var items []utils.Items
	for _, cont := range *conts {
		items = append(items, utils.Items{
			Name:    cont.Names[0],
			Created: humanize.Time(time.Unix(cont.Created, 0)),
			Size:    humanize.Bytes(uint64(cont.SizeRootFs)),
			Status:  cont.Status,
			ID:      cont.ID[:10],
		})
	}
	spew.Dump(items)
	return c.Render("containerlist", utils.ViewData{Name: "test", Items: items})
}

func HandlerContainerPrune(c *fiber.Ctx) error {
	pruned, prunedErr := utils.Visor.ContainerPrune()
	if prunedErr != nil {
		log.Println(prunedErr)
		return c.SendStatus(500)
	}
	_ = c.SendStatus(200)
	return c.SendString(fmt.Sprintf("freed %s", humanize.Bytes(pruned)))
}
