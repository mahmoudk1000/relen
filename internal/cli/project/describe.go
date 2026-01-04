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

func NewDescribeCommand() *cobra.Command {
	var queries *database.Queries

	describe := &cobra.Command{
		Use:     "describe <name>",
		Aliases: []string{"desc"},
		Short:   "Describe a project in detail",
		PreRun: func(cmd *cobra.Command, args []string) {
			queries = db.Get()
		},
	}

	flags := describe.Flags()
	flags.Bool("json", false, "output in JSON format")

	describe.RunE = func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		ctx := cmd.Context()

		jsonFlag, _ := flags.GetBool("json")

		var fmtP string
		p, err := describeProject(ctx, args[0], queries)
		if err != nil {
			return err
		}

		switch {
		case jsonFlag:
			fmtP, err = utils.FormatJSON(p)
			if err != nil {
				return err
			}
		default:
			fmtP, err = utils.Format(p)
			if err != nil {
				return err
			}
		}

		fmt.Println(fmtP)

		return nil
	}

	return describe
}

func describeProject(
	ctx context.Context,
	name string,
	q *database.Queries,
) (models.FProject, error) {
	return models.FProject{}, nil
}
