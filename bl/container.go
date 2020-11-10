package bl

import (
	"errors"
	"fmt"
	"github.com/dustin/go-humanize"
	"log"
	"regexp"
	"time"
	"visor/utils"
)

func ContainerStatProcess(contID string) (*utils.Stats, error) {
	Stat, staterr := utils.Visor.ContainerGetStat(contID)
	if staterr != nil {
		log.Println(staterr)
		return nil, staterr
	}
	if Stat.MemoryStats.Usage > 0 {
		//spew.Dump(Stat)
		usedMemory := uint64(Stat.MemoryStats.Usage - Stat.MemoryStats.Stats.Cache)
		cpuDelta := Stat.CPUStats.CPUUsage.TotalUsage
		systemCPUtDelta := Stat.CPUStats.SystemCPUUsage - Stat.PrecpuStats.SystemCPUUsage
		numberCPUs := Stat.CPUStats.OnlineCpus
		cpuUsage := (float32(cpuDelta) / float32(systemCPUtDelta)) * float32(numberCPUs) * 100.0
		Item := utils.Stats{
			ID:   Stat.ID,
			Name: Stat.Name,
			CPU:  fmt.Sprintf("%f%", cpuUsage),
			RAM:  humanize.Bytes(usedMemory),
			Network: fmt.Sprintf("RX: %s | TX: %s",
				humanize.Bytes(uint64(Stat.Networks.Eth0.RxBytes)),
				humanize.Bytes(uint64(Stat.Networks.Eth0.TxBytes))),
		}
		return &Item, nil
	}
	return nil, errors.New("unknown error")
}

func ContainerListProcess() (*utils.ViewData, error) {
	conts, contsErr := utils.Visor.ContainerListAll()

	if contsErr != nil {
		return nil, contsErr
	}
	var items []utils.Items
	for _, cont := range *conts {
		paused, _ := regexp.MatchString("Paused", cont.Status)
		exited, _ := regexp.MatchString("Exited", cont.Status)
		items = append(items, utils.Items{
			Name:    cont.Names[0],
			Image:   cont.Image,
			Command: cont.Command,
			Created: humanize.Time(time.Unix(cont.Created, 0)),
			Status:  cont.Status,
			Ports:   utils.PortsToString(cont.Ports),
			Paused:  paused,
			Exited:  exited,
			ID:      cont.ID[:10],
		})
	}
	return &utils.ViewData{Name: "test", Items: items}, nil
}
