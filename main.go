package main

import (
	"os"

	"github.com/cloudronics/fileoccurencebeat/cmd"

	_ "github.com/cloudronics/fileoccurencebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
