package project

type InsertProjectRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	AccessLevel int8   `json:"accessLevel" form:"accessLevel" validate:"gte=0,required"`
}
