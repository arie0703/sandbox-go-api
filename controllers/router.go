package controller

import (
	"sandbox-go-api/database"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO!",
		})
	})

	r.GET("user", func(c *gin.Context) {
		db := database.GetDB()

		var user database.User
		// SELECT * FROM users ORDER BY id LIMIT 1;
		db.First(&user)

		c.JSON(200, gin.H{
			"id":   user.ID,
			"name": user.Name,
		})
	})
	return r
}
