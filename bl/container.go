package bl

import (
	"errors"
	"fmt"
	"github.com/dustin/go-humanize"
	"log"
	"time"
	. "visor/utils"
)

func ContainerStatProcess(contID string) (*ViewStat, error) {
	Stat, staterr := Visor.ContainerGetStat(contID)
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
		return &ViewStat, nil
	}
	return nil, errors.New("unknown error")
}

func ContainerListProcess() (*ViewData, error) {
	conts, contsErr := Visor.ContainerListAll()
	if contsErr != nil {
		return nil, contsErr
	}
	var items []Items
	for _, cont := range *conts {
		items = append(items, Items{
			Name:    cont.Names[0],
			Created: humanize.Time(time.Unix(cont.Created, 0)),
			Size:    humanize.Bytes(uint64(cont.SizeRootFs)),
			Status:  cont.Status,
			ID:      cont.ID[:10],
		})
	}
	return &ViewData{Name: "test", Items: items}, nil
}
