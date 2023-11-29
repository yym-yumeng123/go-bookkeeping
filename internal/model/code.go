package model

import "gorm.io/gorm"

type ValidationCode struct {
	gorm.Model
	Code   string `gorm:"not null;"`
	Email  string `gorm:"not null;"`
	UsedAt int
}
