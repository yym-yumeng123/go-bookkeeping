package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID    int
	Email string
	Phone int
	*gorm.Model
}

// 连接数据库
func GormConnect() {
	dsn := "root:YYm1994@tcp(127.0.0.1:3306)/bookkeeping?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	DB = db
	if err != nil {
		panic("failed to connect database")
	}
}

// 创建一个表
func CtreateUserTable() {
	err := DB.Migrator().CreateTable(&User{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("createTables success")
}

// 迁移数据 对数据字段修改, 能增加字段, 不能删除字段
func Migrate() {
	DB.AutoMigrate(&User{})
}

func Curd() {
	// 创建一个 user
	user := &User{Email: "2@qq.com"}
	DB.Create(user)

	// 修改
	user.Phone = 123456787
	DB.Save(user)
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}
