package test

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
	"visor/dock"
)

func Test_ListAllContainer(t *testing.T) {
	v := dock.NewVisor()
	spew.Dump(v.ListAllContainer())
}

func Test_FindContainer(t *testing.T) {
	v := dock.NewVisor()
	spew.Dump(v.FindContainer())
}

func Test_StartContainer(t *testing.T) {
	v := dock.NewVisor()
	spew.Dump(v.StartContainer(""))
}

func Test_PauseContainer(t *testing.T) {
	v := dock.NewVisor()
	spew.Dump(v.PauseContainer(""))
}
