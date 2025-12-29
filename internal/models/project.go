package models

import "time"

type Projects struct {
	Project []Project `json:"projects"`
}

type Project struct {
	Name        string    `json:"name"`
	Link        string    `json:"link,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
