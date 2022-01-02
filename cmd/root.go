package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "git-utils",
		Short: "Utils to simplify git workflows",
		Long:  `git-utils is a set of cli-commands to set/update configs and hooks to modify commit messages`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newUserstoryCmd())
}
