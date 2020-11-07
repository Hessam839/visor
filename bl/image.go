package bl

import (
	"github.com/dustin/go-humanize"
	"log"
	"time"
	. "visor/utils"
)

func ImageListProcess() (*ViewData, error) {
	imgs, imgsErr := Visor.ImageGetAll()
	if imgsErr != nil {
		return nil, imgsErr
	}
	var images []Items
	for _, img := range *imgs {
		images = append(images, Items{
			Name:    img.RepoTags[0],
			Created: humanize.Time(time.Unix(img.Created, 0)),
			Size:    humanize.Bytes(uint64(img.VirtualSize)),
		})
	}
	return &ViewData{Name: "test", Items: images}, nil
}
