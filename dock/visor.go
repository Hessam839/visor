package dock

import (
	"context"
	"github.com/docker/docker/client"
	"log"
	"time"
)

type Visor struct {
	Client  *client.Client
	TimeOut time.Duration
	Ctx     context.Context
}

func NewVisor() *Visor {
	client, err := client.NewClientWithOpts(client.WithHost("tcp://172.17.0.1:2375"), client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	return &Visor{
		Client:  client,
		TimeOut: 30 * time.Second,
		Ctx:     context.Background(),
	}
}
