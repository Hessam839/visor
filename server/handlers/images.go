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

func HandlerGetAllImages(c *fiber.Ctx) error {
	imgs, imgsErr := utils.Visor.ImageGetAll()
	if imgsErr != nil {
		log.Fatal(imgsErr)
	}
	var images []utils.Items
	for _, img := range *imgs {
		images = append(images, utils.Items{
			Name:    img.RepoTags[0],
			Created: humanize.Time(time.Unix(img.Created, 0)),
			Size:    humanize.Bytes(uint64(img.VirtualSize)),
		})
	}
	spew.Dump(images)
	return c.Render("imagelist", utils.ViewData{Name: "test", Items: images})
}

func HandlerImagePrune(c *fiber.Ctx) error {
	pruned, prounedErr := utils.Visor.ImagePrune()
	if prounedErr != nil {
		log.Println(prounedErr)
		return c.SendStatus(500)
	}
	_ = c.SendStatus(200)
	return c.SendString(fmt.Sprintf("freed %s", humanize.Bytes(pruned)))
}
