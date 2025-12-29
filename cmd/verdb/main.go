/*
Copyright © 2026 mahmoudk1000 <mahmoudk1000@gmail.com>
*/
package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/mahmoudk1000/verdb/internal/cli/application"
	"github.com/mahmoudk1000/verdb/internal/cli/project"
)

func main() {
	var verdb = &cobra.Command{
		Use:   "verdb",
		Short: "A serious, well-scoped versioning tool.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}
			return nil
		},
	}

	verdb.AddCommand(project.NewProjectCommand())
	verdb.AddCommand(application.NewApplicationCommand())

	if err := verdb.Execute(); err != nil {
		os.Exit(1)
	}
}
