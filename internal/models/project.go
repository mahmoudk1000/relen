package models

import "time"

type Projects struct {
	Project map[string]Project `json:"projects"`
}

type Project struct {
	Link        string    `json:"link,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
