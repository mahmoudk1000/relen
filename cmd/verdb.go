/*
Copyright © 2026 mahmoudk1000 <mahmoudk1000@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/mahmoudk1000/verdb/internal/cli/application"
	"github.com/mahmoudk1000/verdb/internal/cli/project"
)

// rootCmd represents the base command when called without any subcommands
var verdb = &cobra.Command{
	Use:   "verdb",
	Short: "A CLI for managing project releases composed of multiple applications",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := verdb.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	verdb.AddCommand(project.NewCommand())
	verdb.AddCommand(application.NewCommand())
}
