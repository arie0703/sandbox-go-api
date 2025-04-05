package controller

import (
	"fmt"
	"net/http"
	"sandbox-go-api/crud-api/database"
	model "sandbox-go-api/crud-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateRecipeRequest struct {
	Recipe      model.Recipe
	Ingredients []model.Ingredient
	Procedures  []model.Procedure
}

func GetAllRecipe(c *gin.Context) {
	Recipes := model.GetAllRecipe()

	c.JSON(200, gin.H{
		"Recipes": Recipes,
	})
}

func GetOneRecipe(c *gin.Context) {
	// パスパラメータからid取得
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOneUser(id)

	c.JSON(200, gin.H{
		"id":   data.ID,
		"name": data.Name,
	})
}

func CreateRecipe(c *gin.Context) {
	var request CreateRecipeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	recipe := model.Recipe{
		Title:       request.Recipe.Title,
		Description: request.Recipe.Description,
		UserID:      request.Recipe.UserID,
	}

	var ingredients []model.Ingredient
	var procedures []model.Procedure

	db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&recipe).Error; err != nil {
			return err
		}

		fmt.Println(&recipe.ID)

		// recipe.IDに紐づくingredient, procedureのレコード作成
		for _, v := range request.Ingredients {
			v.RecipeID = recipe.ID
			ingredients = append(ingredients, v)
		}
		for _, v := range request.Procedures {
			v.RecipeID = recipe.ID
			procedures = append(procedures, v)
		}

		if err := tx.Create(&ingredients).Error; err != nil {
			return err
		}

		if err := tx.Create(&procedures).Error; err != nil {
			return err
		}

		return nil
	})

	c.JSON(200, gin.H{
		"message":     "Created Recipe",
		"id":          recipe.ID,
		"title":       recipe.Title,
		"ingredients": ingredients,
		"procedures":  procedures,
	})
}
