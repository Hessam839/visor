package handlers

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
	. "visor/server"
	. "visor/utils"
)

func HandlerGetAllImages(c *fiber.Ctx) error {
	imgs, imgsErr := Visor.GetAllImages()
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
}
