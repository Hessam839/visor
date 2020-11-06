package dock

import (
	"bytes"
	"encoding/json"
	"log"
	"visor/dock/structs"
)

func (v *Visor) ContainerGetStat(ContID string) (*structs.ContainerStat, error) {
	stat, staterr := v.Client.ContainerStats(v.Ctx, ContID, false)
	if staterr != nil {
		return nil, staterr
	}

	buff := new(bytes.Buffer)
	_, buffErr := buff.ReadFrom(stat.Body)
	if buffErr != nil {
		return nil, buffErr
	}
	jstr := buff.String()
	//var data map[string]interface{}
	var data structs.ContainerStat
	jerr := json.Unmarshal([]byte(jstr), &data)
	if jerr != nil {
		log.Fatal(jerr)
	}

	return &data, nil
}
