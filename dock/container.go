package dock

import (
	//docker "github.com/docker/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	network2 "docker.io/go-docker/api/types/network"
)

func (v *Visor) ListAllContainer() (*[]types.Container, error) {
	containers, err := v.Client.ContainerList(v.Ctx, types.ContainerListOptions{Size: true})
	if err != nil {
		return nil, err
	}
	return &containers, nil
}

func (v *Visor) FindContainer() (*[]types.Container, error) {

	return nil, nil
}

func (v *Visor) StartContainer(ContID string) error {
	err := v.Client.ContainerStart(v.Ctx, ContID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (v *Visor) PauseContainer(ContID string) error {
	err := v.Client.ContainerPause(v.Ctx, ContID)
	if err != nil {
		return err
	}
	return nil
}

func (v *Visor) StopContainer(ContID string) error {
	err := v.Client.ContainerStop(v.Ctx, ContID, &v.TimeOut)
	if err != nil {
		return err
	}
	return nil
}

func (v *Visor) KillContainer(ContID string) error {
	err := v.Client.ContainerKill(v.Ctx, ContID, "Kill")
	if err != nil {
		return err
	}
	return nil
}

func (v *Visor) CreateContainer(ContName string) (string, error) {
	config := container.Config{}
	host := container.HostConfig{}
	network := network2.NetworkingConfig{}
	cont, err := v.Client.ContainerCreate(v.Ctx, &config, &host, &network, ContName)
	if err != nil {
		return "", err
	}

	return cont.ID, nil
}
