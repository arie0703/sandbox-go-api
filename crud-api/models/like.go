package model

import (
	"sandbox-go-api/crud-api/database"
	"time"
)

type Like struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RecipeID  uint
	Recipe    Recipe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func GetAllLike() (datas []Like) {
	db := database.GetDB()
	result := db.Find(&datas)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOneLike(id int) (data Like) {
	db := database.GetDB()
	result := db.First(&data, id)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func CreateLike(data *Like) {
	db := database.GetDB()
	result := db.Create(data)
	if result.Error != nil {
		panic(result.Error)
	}
}

func DeleteLike(data *Like) {
	db := database.GetDB()
	result := db.Delete(data)
	if result.Error != nil {
		panic(result.Error)
	}
}
