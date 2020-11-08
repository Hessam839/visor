package bl

import (
	"github.com/docker/docker/api/types/registry"
	"github.com/dustin/go-humanize"
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
			ID:      img.ID[7:17],
		})

	}
	return &ViewData{Name: "test", Items: images}, nil
}

func ImageRemove(imgID string, forced bool) error {
	err := Visor.ImageRemove(imgID, forced)
	if err != nil {
		return err
	}
	return nil
}

func ImageSearch(imgName string) (*[]registry.SearchResult, error) {
	result, resultErr := Visor.ImageSearch(imgName)
	if resultErr != nil {
		return nil, resultErr
	}
	return result, nil
}
