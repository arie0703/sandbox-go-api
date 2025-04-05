package model

import (
	"sandbox-go-api/database"
)

type Ingredient struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Amount   int
	Unit     string
	RecipeID uint
	Recipe   Recipe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func GetAllIngredient() (datas []Ingredient) {
	db := database.GetDB()
	// SELECT * FROM Ingredients ORDER BY id LIMIT 1;
	result := db.Find(&datas)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOneIngredient(id int) (data Ingredient) {
	db := database.GetDB()
	result := db.First(&data, id)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func CreateIngredients(data *[]Ingredient) {
	db := database.GetDB()
	result := db.Create(data)
	if result.Error != nil {
		panic(result.Error)
	}
}

func UpdateIngredient(data *Ingredient) {
	db := database.GetDB()
	result := db.Save(data)
	if result.Error != nil {
		panic(result.Error)
	}
}

func DeleteIngredient(data *Ingredient) {
	db := database.GetDB()
	result := db.Delete(data)
	if result.Error != nil {
		panic(result.Error)
	}
}
