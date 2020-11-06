package dock

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func (v *Visor) ImageGetAll() (*[]types.ImageSummary, error) {
	imagesList, imageErr := v.Client.ImageList(v.Ctx, types.ImageListOptions{})
	if imageErr != nil {
		return nil, imageErr
	}
	return &imagesList, nil
}

func (v *Visor) ImagePrune() (uint64, error) {
	preport, preportErr := v.Client.ImagesPrune(v.Ctx, filters.Args{})
	if preportErr != nil {
		return 0, preportErr
	}
	return preport.SpaceReclaimed, nil
}

func (v *Visor) Image(imageName string) {
	v.Client.ImagePull(v.Ctx, imageName, types.ImagePullOptions{})
}
