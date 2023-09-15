package main

import (
	"os"

	"github.com/aethiopicuschan/kaban/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
