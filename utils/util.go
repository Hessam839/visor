package utils

import "visor/dock"

var Visor *dock.Visor

func init() {
	Visor = dock.NewVisor()
}
