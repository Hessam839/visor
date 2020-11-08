package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
	"visor/bl"
	"visor/utils"
)

func HandlerContainerStats(c *fiber.Ctx) error {
	start := time.Now()
	id := c.Params("id")
	stat, statErr := bl.ContainerStatProcess(id)
	if statErr != nil {
		_ = c.SendStatus(500)
		return c.SendString(statErr.Error())
	}
	log.Printf("response time %s", time.Now().Sub(start))
	return c.Render("containerStat", utils.ViewStat{Name: "stat", Item: []utils.Stats{*stat}})

}
