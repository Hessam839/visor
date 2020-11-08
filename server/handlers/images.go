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

func HandlerImageRemove(c *fiber.Ctx) error {
	id := c.Params("id")
	err := bl.ImageRemove(id, true)
	if err != nil {
		c.Status(500)
		return c.SendString(err.Error())
	}
	c.Status(200)
	return c.SendString("success")
	//return c.Render("imageRemove",)
}

func HandlerImageSearch(c *fiber.Ctx) error {
	iname := c.Params("iname")
	results, resultsErr := bl.ImageSearch(iname)
	if resultsErr != nil {
		c.Status(500)
		return c.SendString(resultsErr.Error())
	}
	var search []utils.Search
	for _, result := range *results {
		search = append(search, utils.Search{
			Name:        result.Name,
			Description: result.Description,
			IsOfficial:  result.IsOfficial,
			Stars:       result.StarCount,
		})
	}
	return c.Render("searchList", utils.ViewSearch{Name: "search", Items: search})
}
