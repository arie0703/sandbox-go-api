package controller

import (
	"fmt"
	"net/http"
	model "sandbox-go-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllSnack(c *gin.Context) {
	snacks := model.GetAllSnack()

	c.JSON(200, gin.H{
		"snacks": snacks,
	})
}

func GetOneSnack(c *gin.Context) {
	// パスパラメータからid取得
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOneUser(id)

	c.JSON(200, gin.H{
		"id":   data.ID,
		"name": data.Name,
	})
}

func CreateSnack(c *gin.Context) {
	var snackJson model.Snack
	if err := c.ShouldBindJSON(&snackJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(snackJson.Name)

	data := model.Snack{Name: snackJson.Name}
	data.CreateSnack()

	c.JSON(200, gin.H{
		"message": "Created snack.",
		"id":      data.ID,
		"name":    data.Name,
	})
}

func EditSnack(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOneSnack(id)

	var snackJson model.Snack
	if err := c.ShouldBindJSON(&snackJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.Name = snackJson.Name
	data.UpdateSnack()

	c.JSON(200, gin.H{
		"message": "Updated snack.",
		"id":      data.ID,
		"name":    data.Name,
	})
}

func DeleteSnack(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOneSnack(id)
	data.DeleteSnack()

	c.JSON(200, gin.H{
		"message": "Deleted snack.",
		"id":      data.ID,
		"name":    data.Name,
	})
}
