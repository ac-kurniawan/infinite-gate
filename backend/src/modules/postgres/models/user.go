package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `json:"username" gorm:"not null; varchar(255)"`
	Email    string    `json:"email" gorm:"not null; varchar(255)"`
	Password string    `json:"password" gorm:"not null; text"`
	Projects []Project `gorm:"many2many:users_projects"`
}
