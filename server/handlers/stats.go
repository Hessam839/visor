package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
	"visor/bl"
	. "visor/utils"
)

func HandlerContainerStats(c *fiber.Ctx) error {
	start := time.Now()
	id := c.Params("id")
	viewStat, viewStatErr := bl.ContainerStatProcess(id)
	if viewStatErr != nil {
		_ = c.SendStatus(500)
		return c.SendString(viewStatErr.Error())
	}
	log.Printf("response time %s", time.Now().Sub(start))
	return c.Render("containerStat", viewStat)

}
