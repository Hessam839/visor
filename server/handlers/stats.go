package handlers

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
	. "visor/utils"
)

func HandlerContainerStats(c *fiber.Ctx) error {
	start := time.Now()
	id := c.Params("id")
	Stat, staterr := Visor.ContainerGetStat(id)
	if staterr != nil {
		log.Println(staterr)
		return c.SendString(staterr.Error())
	}
	if Stat.MemoryStats.Usage > 0 {
		//spew.Dump(Stat)
		usedMemory := uint64(Stat.MemoryStats.Usage - Stat.MemoryStats.Stats.Cache)
		cpuDelta := Stat.CPUStats.CPUUsage.TotalUsage
		systemCPUtDelta := Stat.CPUStats.SystemCPUUsage - Stat.PrecpuStats.SystemCPUUsage
		numberCPUs := Stat.CPUStats.OnlineCpus
		cpuUsage := (float32(cpuDelta) / float32(systemCPUtDelta)) * float32(numberCPUs) * 100.0
		ViewStat := ViewStat{
			Name: "test",
			Item: []Stats{Stats{
				ID:   Stat.ID,
				Name: Stat.Name,
				CPU:  fmt.Sprintf("%f%", cpuUsage),
				RAM:  humanize.Bytes(usedMemory),
				Network: fmt.Sprintf("RX: %s | TX: %s",
					humanize.Bytes(uint64(Stat.Networks.Eth0.RxBytes)),
					humanize.Bytes(uint64(Stat.Networks.Eth0.TxBytes))),
			}},
		}
		log.Printf("response time %s", time.Now().Sub(start))
		return c.Render("containerstat", ViewStat)
	}
	return c.SendStatus(500)
}
