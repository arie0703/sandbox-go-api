package controller

import (
	"fmt"
	"net/http"
	model "sandbox-go-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	users := model.GetAllUser()

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetOneUser(c *gin.Context) {
	// パスパラメータからid取得
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOneUser(id)

	c.JSON(200, gin.H{
		"id":   data.ID,
		"name": data.Name,
	})
}

func CreateUser(c *gin.Context) {
	var userJson model.User
	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(userJson.Name)

	data := model.User{Name: userJson.Name}
	data.CreateUser()

	c.JSON(200, gin.H{
		"message": "Deleted snack.",
		"id":      data.ID,
		"name":    data.Name,
	})
}

func EditUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOneUser(id)

	var snackJson model.Snack
	if err := c.ShouldBindJSON(&snackJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.Name = snackJson.Name
	data.UpdateUser()

	c.JSON(200, gin.H{
		"message": "Deleted snack.",
		"id":      data.ID,
		"name":    data.Name,
	})
}

func DeleteUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOneUser(id)
	data.DeleteUser()

	c.JSON(200, gin.H{
		"message": "Deleted snack.",
		"id":      data.ID,
		"name":    data.Name,
	})
}
