package model

type User struct {
	ID    int `gorm:"PrimaryKey;AutoIncrement"`
	Email string
}
