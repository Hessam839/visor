package handlers

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"log"
	"visor/bl"
	"visor/utils"
)

func HandlerGetAllImages(c *fiber.Ctx) error {
	viewData, viewDataErr := bl.ImageListProcess()
	if viewDataErr != nil {
		log.Println(viewDataErr)
		_ = c.SendStatus(500)
		return c.SendString(viewDataErr.Error())
	}
	return c.Render("imageList", viewData)
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
