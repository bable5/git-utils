package main

import (
	"os"

	"github.com/bable5/git-utils/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
