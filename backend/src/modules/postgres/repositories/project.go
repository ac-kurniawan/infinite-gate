package repositories

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"infinitegate/src/modules/postgres/models"
	"infinitegate/src/services/project"
)

// PGProjectRepo object
type PGProjectRepo struct {
	record *gorm.DB
}

func (pg *PGProjectRepo) FindProjectsByUserID(UserID int) ([]project.Project, error) {
	var projectRecord []models.Project
	var users []models.User
	errUser := pg.record.Find(&users, UserID).Error

	if errUser != nil {
		log.Errorf("[FindProjectByID] - %s", errUser.Error())
		return nil, errUser
	}

	err := pg.record.Model(&users).Association("Projects").Find(&projectRecord)

	if err != nil {
		log.Errorf("[FindProjectByID] - %s", err.Error())
		return nil, err
	}

	var projectsBuffer []project.Project

	for _, v := range projectRecord {
		projectsBuffer = append(projectsBuffer, project.Project{
			ID:          int(v.ID),
			Name:        v.Name,
			AccessLevel: v.AccessLevel,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return projectsBuffer, nil
}

func (pg *PGProjectRepo) InsertProjectByUserID(UserID int, project2 project.Project) (*project.Project, error) {
	var user models.User

	errUsr := pg.record.Find(&user, UserID).Error

	if errUsr != nil {
		return nil, errUsr
	}

	p := models.Project{
		Name:        project2.Name,
		AccessLevel: project2.AccessLevel,
	}

	result := pg.record.Create(&p)

	pg.record.Create(&models.UsersProject{
		UserID:      int(user.ID),
		ProjectID:   int(p.ID),
		AccessLevel: 1, // add default access level for project creator
	})

	return &project.Project{
		ID:          int(p.ID),
		Name:        p.Name,
		AccessLevel: p.AccessLevel,
		CreatedAt:   p.CreatedAt,
	}, result.Error
}

// FindProjectByID methode of PGProjectRepo
func (pg *PGProjectRepo) FindProjectByID(UserID int, ID int) (*project.Project, error) {
	var users []models.User
	errUser := pg.record.Find(&users, UserID).Error

	if errUser != nil {
		log.Errorf("[FindProjectByID] - %s", errUser.Error())
		return nil, errUser
	}

	var projectRecord models.Project

	err := pg.record.Model(&users).Association("Projects").Find(&projectRecord, ID)

	if err != nil {
		log.Errorf("[FindProjectByID] - %s", err.Error())
		return nil, err
	}

	return &project.Project{
		ID:          int(projectRecord.ID),
		Name:        projectRecord.Name,
		AccessLevel: projectRecord.AccessLevel,
		CreatedAt:   projectRecord.CreatedAt,
		UpdatedAt:   projectRecord.UpdatedAt,
	}, nil
}

func NewProjectPG(db *gorm.DB) PGProjectRepo {
	repo := PGProjectRepo{
		db,
	}
	return repo
}
