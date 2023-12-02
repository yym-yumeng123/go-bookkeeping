package model

import "time"

type Item struct {
	ID         int   `gorm:"PrimaryKey;AutoIncrement"`
	Amount     int32 `gorm:"not null"`
	Kind       int   `gorm:"not null"`
	HappenedAt time.Time
	UserId     int
}
