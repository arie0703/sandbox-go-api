package model

import (
	"sandbox-go-api/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func GetAllUser() (datas []User) {
	db := database.GetDB()
	// SELECT * FROM users ORDER BY id LIMIT 1;
	result := db.Find(&datas)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOneUser(id int) (data User) {
	db := database.GetDB()
	// SELECT * FROM users ORDER BY id LIMIT 1;
	result := db.First(&data, id)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (u *User) CreateUser() {
	db := database.GetDB()
	result := db.Create(u)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u *User) UpdateUser() {
	db := database.GetDB()
	result := db.Save(u)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u *User) DeleteUser() {
	db := database.GetDB()
	result := db.Delete(u)
	if result.Error != nil {
		panic(result.Error)
	}
}
