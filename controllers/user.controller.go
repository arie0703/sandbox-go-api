package controller

import (
	"fmt"
	"net/http"
	model "sandbox-go-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	users := model.GetAll()

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetOneUser(c *gin.Context) {
	user := model.GetOne()

	c.JSON(200, gin.H{
		"id":   user.ID,
		"name": user.Name,
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
	data.Create()
}
