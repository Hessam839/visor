package dock

import (
	"context"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Visor struct {
	Client  *client.Client
	Ctx     context.Context
	SysInfo *types.Info
	TimeOut time.Duration
}

func NewVisor() *Visor {
	dockerClient, dockerClientErr := client.NewClientWithOpts(client.WithHost("tcp://172.17.0.1:2375"), client.WithAPIVersionNegotiation())
	if dockerClientErr != nil {
		log.Fatal(dockerClientErr)
	}
	v := &Visor{
		Client:  dockerClient,
		Ctx:     context.Background(),
		TimeOut: 30 * time.Second,
	}

	sysInfo, sysInfoErr := v.Client.Info(v.Ctx)
	if sysInfoErr != nil {
		log.Fatal(sysInfoErr)
	}
	v.SysInfo = &sysInfo
	return v
}
