package project

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mahmoudk1000/relen/internal/database"
	"github.com/mahmoudk1000/relen/internal/db"
	"github.com/mahmoudk1000/relen/internal/models"
	"github.com/mahmoudk1000/relen/internal/utils"
)

func NewListCommand() *cobra.Command {
	var queries *database.Queries

	list := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all projects",
		PreRun: func(cmd *cobra.Command, args []string) {
			queries = db.Get()
		},
	}

	flags := list.Flags()
	flags.Bool("json", false, "Output in JSON format")

	list.RunE = func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		ctx := cmd.Context()

		jsonFlag, _ := flags.GetBool("json")

		ps, err := listProjects(ctx, queries)
		if err != nil {
			return err
		}

		var fmtP string

		switch {
		case jsonFlag:
			fmtP, err = utils.FormatJSON(models.ToProjects(ps))
			if err != nil {
				return err
			}
		default:
			fmtP, err = utils.Format(models.ToProjects(ps))
			if err != nil {
				return err
			}
		}

		fmt.Println(fmtP)
		return nil
	}

	return list
}

func listProjects(ctx context.Context, q *database.Queries) ([]database.Project, error) {
	ps, err := q.ListAllProjects(ctx)
	if err != nil {
		return nil, err
	}

	return ps, nil
}
