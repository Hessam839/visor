package test

import (
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
	"visor/dock"
)

func Test_GetContainerStat(t *testing.T) {
	v := dock.NewVisor()
	s, err := v.GetContainerStat("04dca0882c")
	//s, err := v.GetContainerStat("04dca0882c9e47bdf2c30d1bef55321e5ad2b576c73f369c01a46b5f954f48ce")
	if err != nil {
		log.Fatal(err)
	}
	//spew.Dump(s)
	//buf := new(bytes.Buffer)
	//_, _ = buf.ReadFrom(s.Body)
	//jstr := buf.String()
	////var data map[string]interface{}
	//var data structs.ContainerStat
	//jerr := json.Unmarshal([]byte(jstr), &data)
	//if jerr != nil {
	//	log.Fatal(jerr)
	//}
	////spew.Dump(data)
	////d := data["memory_stats"].(map[string]interface{})
	spew.Dump(s)
}
