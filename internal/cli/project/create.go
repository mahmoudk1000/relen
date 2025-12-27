package project

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/mahmoudk1000/verdb/internal/models"
	"github.com/mahmoudk1000/verdb/internal/utils"
)

func createCommand() *cobra.Command {
	var (
		link        string
		description string
		configPath  string
	)

	create := &cobra.Command{
		Use:     "create <name>",
		Aliases: []string{"c", "new"},
		Short:   "add a new application to the project",
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			path, err := utils.ProjectConfigFilePath()
			if err != nil {
				return err
			}
			configPath = path
			return nil
		},
	}

	flags := create.Flags()
	flags.StringVarP(&link, "link", "l", "", "link to the project")
	flags.StringVarP(&description, "description", "d", "", "description of the application")

	create.RunE = func(cmd *cobra.Command, args []string) error {
		err := createJSONProject(configPath, args[0], link, description)
		if err != nil {
			return fmt.Errorf("failed to create project: %w", err)
		}
		return nil
	}

	return create
}

func createJSONProject(path, name, link, desc string) error {
	var p models.Projects
	p.Project = make(map[string]models.Project)

	if f, err := os.Open(path); err == nil {
		defer f.Close()
		decoder := json.NewDecoder(f)
		_ = decoder.Decode(&p)
	}

	p.Project[name] = models.Project{
		Link:        link,
		Description: desc,
		CreatedAt:   time.Now().UTC(),
	}

	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	encoder := json.NewEncoder(f)
	if err := encoder.Encode(p); err != nil {
		return err
	}

	return nil
}
