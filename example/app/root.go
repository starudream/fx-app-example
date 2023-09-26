package main

import (
	"github.com/starudream/go-lib/v2/app"
)

var rootCmd = app.RootCommand(func(c *app.Command) {
	c.Use = "example-app"
})
