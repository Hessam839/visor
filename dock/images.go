package dock

import (
	"github.com/docker/docker/api/types"
)

func (v *Visor) GetAllImages() (*[]types.ImageSummary, error) {
	imagesList, imageErr := v.Client.ImageList(v.Ctx, types.ImageListOptions{})
	if imageErr != nil {
		return nil, imageErr
	}
	return &imagesList, nil
}
