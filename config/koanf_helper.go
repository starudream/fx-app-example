package config

import (
	"strings"

	"github.com/spf13/pflag"

	"github.com/starudream/go-lib/v2/internal/koanf/providers/posflag"
)

func envCB() func(key string, value string) (string, any) {
	return func(key string, value string) (string, any) {
		key = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(key), "__", "."))
		if _, ok := AvailEnvKeys[key]; ok {
			return key, value
		}
		if !strings.HasPrefix(key, EnvPrefix) {
			return "", nil
		}
		return key, value
	}
}

func flagCB(fs *pflag.FlagSet) func(f *pflag.Flag) (string, any) {
	return func(f *pflag.Flag) (string, any) {
		return strings.ReplaceAll(strings.TrimSpace(f.Name), "__", "."), posflag.FlagVal(fs, f)
	}
}
