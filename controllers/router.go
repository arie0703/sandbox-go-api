package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO!",
		})
	})

	r.GET("/user", GetAllUser)
	r.GET("/user/:id", GetOneUser)
	r.POST("/user/create", CreateUser)
	r.POST("/user/edit/:id", EditUser)
	r.POST("/user/delete/:id", DeleteUser)

	r.GET("/snack", GetAllSnack)
	r.GET("/snack/:id", GetOneSnack)
	r.POST("/snack/create", CreateSnack)
	r.POST("/snack/edit/:id", EditSnack)
	r.POST("/snack/delete/:id", DeleteSnack)
	return r
}
