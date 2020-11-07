package handlers

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"log"
	"visor/bl"
	"visor/utils"
)

func HandlerContainerList(c *fiber.Ctx) error {
	viewData, viewDataErr := bl.ContainerListProcess()
	if viewDataErr != nil {
		log.Println(viewDataErr.Error())
		_ = c.SendStatus(500)
		return c.SendString(viewDataErr.Error())
	}
	//spew.Dump(items)
	return c.Render("containerList", viewData)
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
