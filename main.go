package main

import (
	"webhook/configs"
	"webhook/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	configs.ConnectDB()

	routes.AppRoute(app)

	err := app.Run(":" + configs.GetConfig("server.port"))

	if err != nil {
		panic(err)
	}
}
