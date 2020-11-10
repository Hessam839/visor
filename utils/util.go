package utils

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"visor/dock"
)

var Visor *dock.Visor

func init() {
	Visor = dock.NewVisor()
}

func PortsToString(ports []types.Port) string {
	var portStr string
	for _, port := range ports {
		portStr += fmt.Sprintf("%s:%d->%d/%s ; ", port.IP, port.PrivatePort, port.PublicPort, port.Type)
	}
	return portStr
}
