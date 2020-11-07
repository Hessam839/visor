package dock

import (
	"bytes"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/registry"
	"io"
	"os"
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

	imagesReader, imageReaderErr := v.Client.ImagePull(v.Ctx, imageName, types.ImagePullOptions{
		All:      false,
		Platform: v.Platform})
	if imageReaderErr != nil {
		return "", imageReaderErr
	}
	defer func() {
		_ = imagesReader.Close()
	}()
	io.Copy(os.Stdout, imagesReader)

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

func (v *Visor) ImageRemove(imageID string, force bool) error {
	_, responseErr := v.Client.ImageRemove(v.Ctx, imageID, types.ImageRemoveOptions{
		Force: force,
	})
	if responseErr != nil {
		return responseErr
	}
	return nil
}

func (v *Visor) ImageInspect(imageID string) (*types.ImageInspect, error) {
	detail, _, detailErr := v.Client.ImageInspectWithRaw(v.Ctx, imageID)
	if detailErr != nil {
		return nil, detailErr
	}
	return &detail, nil
}

//func (v *Visor) ImageSave (imageID string) error{
//	response, responseErr := v.Client.ImageSave(v.Ctx, []string{imageID})
//	if responseErr != nil {
//		return responseErr
//	}
//	io.Copy(os.Stdout, response)
//	return nil
//}
//
//func (v *Visor) ImageLoad() error {
//	v.Client.ImageLoad()
//}
