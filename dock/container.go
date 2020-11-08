package dock

import (
	//docker "github.com/docker/go-docker"
	//"docker.io/go-docker/api/types"
	//"docker.io/go-docker/api/types/container"
	//network2 "docker.io/go-docker/api/types/network"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

func (v *Visor) ContainerListAll() (*[]types.Container, error) {
	containers, err := v.Client.ContainerList(v.Ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	return &containers, nil
}

func (v *Visor) ContainerFind(filter types.ContainerListOptions) (*[]types.Container, error) {
	containers, err := v.Client.ContainerList(v.Ctx, filter)
	if err != nil {
		return nil, err
	}
	return &containers, nil
}

func (v *Visor) ContainerStart(ContID string) error {
	err := v.Client.ContainerStart(v.Ctx, ContID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (v *Visor) ContainerPause(ContID string) error {
	err := v.Client.ContainerPause(v.Ctx, ContID)
	if err != nil {
		return err
	}
	return nil
}

func (v *Visor) ContainerStop(ContID string) error {
	err := v.Client.ContainerStop(v.Ctx, ContID, &v.TimeOut)
	if err != nil {
		return err
	}
	return nil
}

func (v *Visor) ContainerKill(ContID string) error {
	err := v.Client.ContainerKill(v.Ctx, ContID, "Kill")
	if err != nil {
		return err
	}
	return nil
}

//TODO mustbe deployed
func (v *Visor) ContainerCreate(ContName string) (string, error) {
	config := container.Config{}
	host := container.HostConfig{}
	network := network.NetworkingConfig{}
	platform := specs.Platform{}
	cont, err := v.Client.ContainerCreate(v.Ctx, &config, &host, &network, &platform, ContName)
	if err != nil {
		return "", err
	}

	return cont.ID, nil
}

func (v *Visor) ContainerPrune() (uint64, error) {
	preport, preportErr := v.Client.ContainersPrune(v.Ctx, filters.Args{})
	if preportErr != nil {
		return 0, preportErr
	}
	return preport.SpaceReclaimed, nil

}
