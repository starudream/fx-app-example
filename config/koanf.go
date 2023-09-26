package config

import (
	"github.com/knadh/koanf/v2"
	"github.com/spf13/cobra"

	"github.com/starudream/go-lib/v2/codec/json"

	"github.com/starudream/go-lib/v2/internal/koanf/parsers/yaml"
	"github.com/starudream/go-lib/v2/internal/koanf/providers/env"
	"github.com/starudream/go-lib/v2/internal/koanf/providers/file"
	"github.com/starudream/go-lib/v2/internal/koanf/providers/posflag"
	"github.com/starudream/go-lib/v2/internal/log"
)

type Koanf = koanf.Koanf

var (
	EnvPrefix    = "app_"
	AvailEnvKeys = map[string]struct{}{
		"debug": {},
	}
)

func LoadFromCommand(cmd *cobra.Command) *Koanf {
	k := koanf.NewWithConf(koanf.Conf{
		Delim:       ".",
		StrictMerge: true,
	})

	logger := log.X.Named("config")

	// file
	path := configPath(cmd)
	if path != "" {
		err := k.Load(file.Provider(path), yaml.Parser())
		if err != nil {
			logger.Fatalf("load file %s error: %v", path, err)
		} else {
			logger.Infof("load file %s success", path)
		}
	} else {
		paths := defaultPaths()
		for i := 0; i < len(paths); i++ {
			path = paths[i]
			err := k.Load(file.Provider(path), yaml.Parser())
			if err == nil {
				logger.Infof("load file %s success", path)
				break
			}
		}
	}

	// env
	_ = k.Load(env.ProviderWithValue("", ".", envCB()), nil)

	// flag
	fs := cmd.PersistentFlags()
	_ = k.Load(posflag.ProviderWithFlag(fs, ".", k, flagCB(fs)), nil)

	if Debug(k) {
		logger.Info(json.MustMarshalString(k.Raw()))
	}

	return k
}
