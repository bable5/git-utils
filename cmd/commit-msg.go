package cmd

import (
	"fmt"
	"os"
	"text/template"

	gitutils "github.com/bable5/git-utils/git"
	"github.com/spf13/cobra"
)

type templateParams struct {
	US string
}

// Stack overflow solution. https://stackoverflow.com/questions/32551811/read-file-as-template-execute-it-and-write-it-back
func applyTemplate(fileName string, p templateParams) error {
	//open the template, if it exists
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, p)
}

func newCommitMsgCmd() *cobra.Command {
	var gitConfigSection string
	commitMsgCmd := &cobra.Command{
		Use:   "commit-msg",
		Short: "git commit-msg hook",
		Long:  `Treat the commit message as a template and template in the userstory, if set.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			userstory, err := gitutils.GetFromConfig(gitConfigSection)
			if err != nil {
				if _, ok := err.(*gitutils.UserStoryNotFound); ok {
					return nil
				} else {
					return err
				}
			}

			if len(args) < 1 {
				return fmt.Errorf("missing template file")
			}

			return applyTemplate(args[0], templateParams{
				US: userstory,
			})
		},
	}

	commitMsgCmd.PersistentFlags().StringVarP(&gitConfigSection, "git-config-section", "", "bable5.userstory", "Name of the config entry to store the userstory in.")

	return commitMsgCmd
}
