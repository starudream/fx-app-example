package main

import (
	"github.com/starudream/go-lib/v2/log"
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.L().Fatal(err)
	}
}
