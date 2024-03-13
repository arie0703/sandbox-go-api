package model

import (
	"sandbox-go-api/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func GetAll() (datas []User) {
	db := database.GetDB()
	// SELECT * FROM users ORDER BY id LIMIT 1;
	result := db.Find(&datas)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOne() (data User) {
	db := database.GetDB()
	// SELECT * FROM users ORDER BY id LIMIT 1;
	result := db.First(&data)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (u *User) Create() {
	db := database.GetDB()
	result := db.Create(u)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u *User) Update() {
	db := database.GetDB()
	result := db.Save(u)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u *User) Delete() {
	db := database.GetDB()
	result := db.Delete(u)
	if result.Error != nil {
		panic(result.Error)
	}
}
