package application

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewAddCommand() *cobra.Command {
	var (
		link        string
		description string
	)

	add := &cobra.Command{
		Use:     "add <project_name> <application_name>",
		Aliases: []string{"c", "new"},
		Short:   "add a new application to the project",
		Args:    cobra.ExactArgs(2),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	flags := add.Flags()
	flags.StringVarP(&link, "link", "l", "", "link to the project")
	flags.StringVarP(&description, "description", "d", "", "description of the application")

	add.RunE = func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Adding application '%s' to project '%s'\n", args[1], args[0])
		return nil
	}

	return add
}
