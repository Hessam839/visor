package dock

import (
	"bytes"
	"encoding/json"
	"log"
	"visor/dock/structs"
)

func (v *Visor) ContainerGetStat(ContID string) (*structs.ContainerStat, error) {
	stat, statErr := v.Client.ContainerStats(v.Ctx, ContID, false)
	if statErr != nil {
		return nil, statErr
	}

	buff := new(bytes.Buffer)
	_, buffErr := buff.ReadFrom(stat.Body)
	if buffErr != nil {
		return nil, buffErr
	}
	defer func() {
		_ = stat.Body.Close()
	}()

	jstr := buff.String()
	var data structs.ContainerStat
	jerr := json.Unmarshal([]byte(jstr), &data)
	if jerr != nil {
		log.Fatal(jerr)
	}

	return &data, nil
}
