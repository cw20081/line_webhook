package routes

import (
	"webhook/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	router.POST("/callback", controllers.SaveChat())

	router.POST("/pushmessage/:userID", controllers.SendChat())

	router.GET("/message/index/:userID", controllers.IndexChat())
}
