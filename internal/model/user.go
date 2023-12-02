package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    int    `gorm:"PrimaryKey;AutoIncrement"`
	Email string `gorm:"unique"`
	Item  []Item // has_many item
}
