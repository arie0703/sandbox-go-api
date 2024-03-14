package model

import (
	"sandbox-go-api/database"
	"time"
)

type Snack struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name"`
}

func GetAllSnack() (datas []Snack) {
	db := database.GetDB()
	// SELECT * FROM Snacks ORDER BY id LIMIT 1;
	result := db.Find(&datas)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOneSnack(id int) (data Snack) {
	db := database.GetDB()
	// SELECT * FROM users ORDER BY id LIMIT 1;
	result := db.First(&data, id)

	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (s *Snack) CreateSnack() {
	db := database.GetDB()
	result := db.Create(s)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (s *Snack) UpdateSnack() {
	db := database.GetDB()
	result := db.Save(s)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (s *Snack) DeleteSnack() {
	db := database.GetDB()
	result := db.Delete(s)
	if result.Error != nil {
		panic(result.Error)
	}
}
