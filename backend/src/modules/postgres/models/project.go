package models

import (
	"gorm.io/gorm"
	"time"
)

type UsersProject struct {
	UserID      int  `gorm:"primaryKey"`
	ProjectID   int  `gorm:"primaryKey"`
	AccessLevel int8 `json:"accessLevel" gorm:"not null"`
	CreatedAt   time.Time
}

type Project struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null; varchar(255)" valid:"required"`
	AccessLevel int8   `json:"accessLevel" gorm:"not null"`
	Users       []User `gorm:"many2many:users_projects"`
}
