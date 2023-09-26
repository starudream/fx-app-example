package main

import (
	"fmt"

	"github.com/starudream/go-lib/v2/app"
)

var actionCmd = app.NewCommand(func(c *app.Command) {
	c.Use = "action"
	c.Short = "action app"
	c.Run = app.Action(actionRun)
})

func init() {
	rootCmd.AddCommand(actionCmd)
}

func actionRun(_ *app.Command, args []string) error {
	fmt.Println("action", args)
	return nil
}
