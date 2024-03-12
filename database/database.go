package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

var (
	db  *gorm.DB
	err error
)

func Init() {
	dsn := "sandbox:develop@/sandbox-go-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func GetDB() *gorm.DB {
	return db
}
