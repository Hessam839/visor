package dock

import (
	"bytes"
	"encoding/json"
	"log"
	"visor/dock/structs"
)

func (v *Visor) GetContainerStat(ContID string) (*structs.ContainerStat, error) {
	stat, staterr := v.Client.ContainerStats(v.Ctx, ContID, false)
	if staterr != nil {
		return nil, staterr
	}

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(stat.Body)
	jstr := buf.String()
	//var data map[string]interface{}
	var data structs.ContainerStat
	jerr := json.Unmarshal([]byte(jstr), &data)
	if jerr != nil {
		log.Fatal(jerr)
	}
	//spew.Dump(data)
	//d := data["memory_stats"].(map[string]interface{})
	//spew.Dump(data)

	return &data, nil
}
