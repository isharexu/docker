package graphdriver

import (
	_ "github.com/docker/docker/daemon/graphdriver/vfs"
	_ "github.com/docker/docker/daemon/graphdriver/windows"
	"github.com/docker/docker/pkg/hcsshim"
)

type WindowsGraphDriver interface {
	Driver
	CopyDiff(id, sourceId string, parentLayerPaths []string) error
	LayerIdsToPaths(ids []string) []string
	Info() hcsshim.DriverInfo
}

var (
	// Slice of drivers that should be used in order
	priority = []string{
		"windowsfilter",
		"windowsdiff",
		"vfs",
	}
)

func GetFSMagic(rootpath string) (FsMagic, error) {
	// Note it is OK to return FsMagicUnsupported on Windows.
	return FsMagicUnsupported, nil
}
