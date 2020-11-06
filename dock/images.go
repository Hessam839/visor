package dock

import (
	"bytes"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/registry"
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

func (v *Visor) ImagePull(imageName string) (string, error) {
	var platform string
	if v.SysInfo.Architecture == "x86_64" {
		platform = v.SysInfo.OSType + "/amd64"
	} else if v.SysInfo.Architecture == "x86" {
		platform = v.SysInfo.OSType + "/386"
	}

	imagesReader, imageReaderErr := v.Client.ImagePull(v.Ctx, imageName, types.ImagePullOptions{
		All:      false,
		Platform: platform})
	if imageReaderErr != nil {
		return "", imageReaderErr
	}
	defer func() {
		_ = imagesReader.Close()
	}()
	buff := bytes.Buffer{}
	_, buffErr := buff.ReadFrom(imagesReader)
	if buffErr != nil {
		return "", buffErr
	}
	jstr := buff.String()
	return jstr, nil
}

func (v *Visor) ImageSearch(imageName string) (*[]registry.SearchResult, error) {
	searchResult, searchResultErr := v.Client.ImageSearch(v.Ctx, imageName, types.ImageSearchOptions{Limit: 10})
	if searchResultErr != nil {
		return nil, searchResultErr
	}
	return &searchResult, nil
}
