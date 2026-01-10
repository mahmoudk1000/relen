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
	flags.Int32P("number", "n", 0, "Number of projects to list (0 for all)")

	list.RunE = func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		ctx := cmd.Context()

		jsonFlag, _ := flags.GetBool("json")
		count, _ := flags.GetInt32("number")

		ps, err := listProjects(ctx, count, queries)
		if err != nil {
			return err
		}

		var fmtP string

		switch {
		case jsonFlag:
			fmtP, err = utils.FormatJSON(ps)
			if err != nil {
				return err
			}
		default:
			fmtP, err = utils.Format(ps)
			if err != nil {
				return err
			}
		}

		fmt.Println(fmtP)
		return nil
	}

	return list
}

func listProjects(ctx context.Context, c int32, q *database.Queries) ([]models.Project, error) {
	var (
		ps  []database.Project
		err error
	)

	if c == 0 {
		ps, err = q.ListAllProjects(ctx)
	} else {
		ps, err = q.ListNProjects(ctx, c)
	}

	if err != nil {
		return nil, err
	}

	return models.ToProjects(ps), nil
}
