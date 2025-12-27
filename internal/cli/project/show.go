package project

import (
	"github.com/spf13/cobra"

	"github.com/mahmoudk1000/verdb/internal/utils"
)

func showCommand() *cobra.Command {
	show := &cobra.Command{
		Use:     "show <name>",
		Aliases: []string{"s"},
		Short:   "show details of a project",
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			_, err := utils.ProjectConfigFilePath()
			return err
		},
	}

	show.RunE = func(cmd *cobra.Command, args []string) error {
		return nil
	}

	return show
}

func showJSONProject(name string) error {
	return nil
}
