package model

import "time"

type Item struct {
	ID         int    `gorm:"PrimaryKey;AutoIncrement"`
	Amount     int32  `gorm:"not null"`
	Kind       string `gorm:"not null"`
	HappenedAt time.Time
	TagIds     int32
}
