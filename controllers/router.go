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
	r.GET("/user/first", GetOneUser)
	r.POST("/user/create", CreateUser)
	return r
}
