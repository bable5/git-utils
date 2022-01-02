package cmd

import (
	"fmt"

	gitutils "github.com/bable5/git-utils/git"
	"github.com/spf13/cobra"
)

var (
	gitConfigSection string
	removeFlag       bool
	userstoryCmd     = &cobra.Command{
		Use:   "userstory [userstory-id]",
		Short: "Set/Update/Clear the userstory",
		Long:  `Manage the 'userstory' git config value.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) >= 1 {
				userstory := args[0]
				return gitutils.AddToConfig(gitConfigSection, userstory)
			} else if removeFlag {
				fmt.Println("removing the config")
				return gitutils.RemoveFromConfig(gitConfigSection)
			} else {

				userstory, err := gitutils.GetFromConfig(gitConfigSection)
				if err != nil {
					return err
				}

				fmt.Println(userstory)

				return nil
			}
		},
	}
)

func init() {
	userstoryCmd.PersistentFlags().StringVarP(&gitConfigSection, "git-config-section", "", "bable5.userstory", "Name of the config entry to store the userstory in.")
	userstoryCmd.Flags().BoolVarP(&removeFlag, "remove", "", false, "clears the userstory config")
}
