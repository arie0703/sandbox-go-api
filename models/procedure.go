package model

import (
	"sandbox-go-api/database"
)

type Procedure struct {
	ID          uint `gorm:"primarykey"`
	OrderNumber int
	Description string
	RecipeID    uint
	Recipe      Recipe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func GetAllProcedure() (datas []Procedure) {
	db := database.GetDB()
	result := db.Find(&datas)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOneProcedure(id int) (data Procedure) {
	db := database.GetDB()
	result := db.First(&data, id)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func CreateProcedures(data *[]Procedure) {
	db := database.GetDB()
	result := db.Create(data)
	if result.Error != nil {
		panic(result.Error)
	}
}

func UpdateProcedure(data *Procedure) {
	db := database.GetDB()
	result := db.Save(data)
	if result.Error != nil {
		panic(result.Error)
	}
}

func DeleteProcedure(data *Procedure) {
	db := database.GetDB()
	result := db.Delete(data)
	if result.Error != nil {
		panic(result.Error)
	}
}
