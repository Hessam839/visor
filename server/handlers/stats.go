package handlers

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"visor/server"
	. "visor/utils"
)

func HandlerContainerStats(c *fiber.Ctx) error {
	id := c.Params("id")
	stat, staterr := Visor.GetContainerStat(id)
	if staterr != nil {
		return staterr
	}
	spew.Dump(stat)
	return c.Render("containerstat", server.Stats{
		ID:   stat.ID,
		Name: stat.Name,
		CPU:  string(stat.CPUStats.CPUUsage.TotalUsage),
		RAM:  humanize.Bytes(uint64(stat.MemoryStats.Usage)),
		Network: fmt.Sprintf("RX: %s | TX: %s",
			humanize.Bytes(uint64(stat.Networks.Eth0.RxBytes)),
			humanize.Bytes(uint64(stat.Networks.Eth0.TxBytes))),
	})
}
