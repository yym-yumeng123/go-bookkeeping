package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name   string `json:"name"`
	UserId int
}

type CreateTagRequest struct {
	Name string `json:"name" building:"required"`
}

func (t Tag) TableName() string {
	return "tags"
}
