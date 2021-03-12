package project

import (
	"time"
)

// IProjectService the repository implementation interface
type IProjectService interface {
	// FindProjectByID Get project by ID
	FindProjectByID(UserID int, ID int) (*Project, error)

	// FindProjectsByUserID get project by user ID
	FindProjectsByUserID(UserID int) ([]Project, error)

	// InsertProjectByUserID Insert new project by user ID
	InsertProjectByUserID(UserID int, project2 Project) (*Project, error)
}

// Service for Project entities
type Service struct {
	projectRepository ProjectRepository
}

func (p Service) FindProjectByID(UserID int, ID int) (*Project, error) {
	result, err := p.projectRepository.FindProjectByID(UserID, ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p Service) FindProjectsByUserID(UserID int) ([]Project, error) {
	results, err := p.projectRepository.FindProjectsByUserID(UserID)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (p Service) InsertProjectByUserID(UserID int, project2 Project) (*Project, error) {
	project := NewProject(
		project2.Name,
		project2.AccessLevel,
		time.Now(),
	)

	identifier, err := p.projectRepository.InsertProjectByUserID(UserID, project)

	if err != nil {
		return nil, err
	}

	return identifier, nil
}

func NewProjectRepositoryImpl(projectRepository ProjectRepository) IProjectService {
	return &Service{
		projectRepository,
	}
}
