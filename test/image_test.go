package test

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
	"visor/dock"
)

func Test_GetAllImages(t *testing.T) {
	v := dock.NewVisor()
	spew.Dump(v.ImageGetAll())
}
