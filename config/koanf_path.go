package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/starudream/go-lib/v2/osx"
)

func configPath(cmd *cobra.Command) (path string) {
	flagConfig := cmd.PersistentFlags().Lookup("config")
	if flagConfig != nil && flagConfig.Changed {
		path = flagConfig.Value.String()
	}
	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}
	if path == "" {
		path = os.Getenv("CONFIG")
	}
	return
}

func defaultPaths() []string {
	exts := []string{
		".yml",
		".yaml",
	}
	pres := []string{
		filepath.Join(osx.ExeDir, osx.ExeName),
		filepath.Join(osx.WorkDir, osx.ExeName),
		filepath.Join(osx.ExeDir, "config"),
		filepath.Join(osx.WorkDir, "config"),
	}

	if dir, _ := os.UserHomeDir(); dir != "" {
		pres = append(
			pres,
			filepath.Join(dir, osx.ExeName),
			filepath.Join(dir, "."+osx.ExeName),
			filepath.Join(dir, ".config", osx.ExeName, "config"),
		)
	}

	if dir, _ := os.UserConfigDir(); dir != "" {
		pres = append(
			pres,
			filepath.Join(dir, osx.ExeName, "config"),
		)
	}

	extsLen := len(exts)
	presLen := len(pres)

	paths := make([]string, presLen*extsLen)

	for i := 0; i < len(pres); i++ {
		for j := 0; j < len(exts); j++ {
			paths[i*extsLen+j] = pres[i] + exts[j]
		}
	}

	return paths
}
