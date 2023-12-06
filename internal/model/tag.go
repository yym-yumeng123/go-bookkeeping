package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

type CreateTagRequest struct {
	Name string `json:"name"`
}

func (t Tag) TableName() string {
	return "tags"
}
