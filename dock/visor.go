package dock

import (
	"context"
	docker "docker.io/go-docker"
	"log"
	"time"
)

type Visor struct {
	Client  *docker.Client
	TimeOut time.Duration
	Ctx     context.Context
}

func NewVisor() *Visor {
	client, err := docker.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}
	return &Visor{
		Client:  client,
		TimeOut: 30 * time.Second,
		Ctx:     context.Background(),
	}
}
