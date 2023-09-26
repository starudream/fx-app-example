package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/starudream/go-lib/v2/app"
	"github.com/starudream/go-lib/v2/codec/yaml"
	"github.com/starudream/go-lib/v2/log"
	"github.com/starudream/go-lib/v2/x"
)

var rootCmd = app.RootCommand(func(c *app.Command) {
	c.Use = "config"
	c.Run = app.Action(configRun)
})

func configRun(_ *app.Command, _ []string) error {
	type Config struct {
		Debug bool       `yaml:"debug" koanf:"debug"`
		Log   log.Config `yaml:"log" koanf:"log"`
	}

	config := Config{
		Debug: false,
		Log: log.Config{
			Console: log.ConfigConsole{
				Level:   "debug",
				NoColor: false,
			},
			File: log.ConfigFile{
				Level:      "info",
				Color:      false,
				Filename:   "bin/app.log",
				MaxSize:    100,
				MaxAge:     180,
				MaxBackups: 10,
				Compress:   true,
			},
		},
	}

	bs, err := yaml.Marshal(
		config,
		yaml.Indent(2),
		yaml.IndentSequence(true),
		yaml.UseLiteralStyleIfMultiline(true),
		yaml.WithComment(yaml.CommentMap{
			"$.debug": []*yaml.Comment{yaml.HeadComment(" debug")},
		}),
	)
	fmt.Println(string(x.Must1(bs, err)))

	path := x.Must1(filepath.Abs("config.yaml"))

	log.L().Infof("write config to %s", path)

	x.Must0(os.WriteFile(path, bs, 0644))

	return nil
}
