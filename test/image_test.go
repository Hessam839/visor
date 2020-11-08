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

func Test_ImagePull(t *testing.T) {
	v := dock.NewVisor()
	spew.Dump(v.ImagePull("redis"))
}

func Test_ImageSearch(t *testing.T) {
	v := dock.NewVisor()
	spew.Dump(v.ImageSearch("hello world"))
}

func Test_ImageRemove(t *testing.T) {
	v := dock.NewVisor()
	spew.Dump(v.ImageRemove("f0b02e9d09", true))
}

func Test_ImageInspect(t *testing.T) {
	v := dock.NewVisor()
	spew.Dump(v.ImageInspect("hello-world"))
}
