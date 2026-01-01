package models

import (
	"github.com/mahmoudk1000/relen/internal/database"
)

type Project struct {
	Name        string `json:"name"`
	Link        string `json:"link,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at"`
}

func DatabaseProjectToProject(p database.Project) Project {
	return Project{
		Name:        p.Name,
		Link:        p.Link.String,
		Description: p.Description.String,
		CreatedAt:   p.CreatedAt.Format("2006-01-02T15:04:05 -07:00:00"),
	}
}
