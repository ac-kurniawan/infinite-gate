package project

type ProjectRepository interface {
	// FindProjectByID Get project by ID
	FindProjectByID(UserID int, ID int) (*Project, error)

	// FindProjectsByUserID get project by user ID
	FindProjectsByUserID(UserID int) ([]Project, error)

	// InsertProjectByUserID Insert new project by user ID
	InsertProjectByUserID(UserID int, project2 Project) (*Project, error)
}
