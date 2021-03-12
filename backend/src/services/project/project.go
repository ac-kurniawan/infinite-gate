package project

import "time"

type Project struct {
	ID          int
	Name        string
	AccessLevel int8
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewProject create new entities project
func NewProject(
	name string,
	accessLevel int8,
	createdAt time.Time) Project {
	return Project{
		Name:        name,
		AccessLevel: accessLevel,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
	}
}
