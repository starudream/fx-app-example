package app

import (
	"sync"

	"github.com/spf13/cobra"

	"github.com/starudream/go-lib/v2/version"
	"github.com/starudream/go-lib/v2/x"
)

type Command = cobra.Command

func NewCommand(fs ...func(c *Command)) *Command {
	cmd := &Command{}
	if len(fs) > 0 && fs[0] != nil {
		fs[0](cmd)
	}
	return cmd
}

var (
	rootCmd     *Command
	rootCmdOnce sync.Once
)

func RootCommand(fs ...func(c *Command)) *Command {
	rootCmdOnce.Do(func() {
		rootCmd = &Command{}
		rootCmd.SetVersionTemplate("{{ print .Version }}")
		rootCmd.Version = x.ToPtr(version.GetVersionInfo()).String()
		rootCmd.PersistentFlags().StringP("config", "c", "", "config file")
		if len(fs) > 0 && fs[0] != nil {
			fs[0](rootCmd)
		}
	})
	return rootCmd
}
