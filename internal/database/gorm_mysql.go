package database

import (
	"bookkeeping/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

var models = []any{&model.User{}, &model.Item{}, &model.Tag{}, &model.ValidationCode{}}

// 连接数据库
func GormConnect() {
	dsn := "root:YYm1994@tcp(127.0.0.1:3306)/bookkeeping?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	fmt.Println("我链接了")
	if err != nil {
		panic("failed to connect database")
	}
}

func CreateTable() {

	for _, model := range models {
		err = DB.Migrator().CreateTable(&model)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("创建表成功")
}

func Migrate() {
	// 自动迁移表字段, 修改 struct 结构体里面的字段, 例: 添加 Phone 字段
	err = DB.AutoMigrate(models...)
	if err != nil {
		log.Println(err)
	}
	log.Println("迁移表成功")
}

func DropColumn() {
	err = DB.Migrator().DropColumn(&model.User{}, "address")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("删除列成功")
}

func Crud() {
	// 创建一个用户 user
	user := model.User{Email: "1@qq.com"}
	tx := DB.Create(&user)
	log.Println(tx.RowsAffected)
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}
