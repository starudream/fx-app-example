package osx

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/starudream/go-lib/v2/x"
)

var (
	WorkDir string
	ExeDir  string
	ExeName string
)

func init() {
	var err error

	WorkDir, err = os.Getwd()
	x.Must0(err)

	exe, err := os.Executable()
	x.Must0(err)

	ExeDir = filepath.Dir(exe)
	ExeName = strings.TrimSuffix(filepath.Base(exe), filepath.Ext(exe))
}
