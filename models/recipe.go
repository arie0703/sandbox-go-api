package model

import (
	"sandbox-go-api/database"
	"time"
)

type Recipe struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description string
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func GetAllRecipe() (datas []Recipe) {
	db := database.GetDB()
	result := db.Find(&datas)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOneRecipe(id int) (data Recipe) {
	db := database.GetDB()
	result := db.First(&data, id)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func CreateRecipe(data *Recipe) {
	db := database.GetDB()
	result := db.Create(data)
	if result.Error != nil {
		panic(result.Error)
	}
}

func UpdateRecipe(data *Recipe) {
	db := database.GetDB()
	result := db.Save(data)
	if result.Error != nil {
		panic(result.Error)
	}
}

func DeleteRecipe(data *Recipe) {
	db := database.GetDB()
	result := db.Delete(data)
	if result.Error != nil {
		panic(result.Error)
	}
}
