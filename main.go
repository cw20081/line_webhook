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

	if err := app.Run(":" + configs.GetConfig("server.port")); err != nil {
		panic(err)
	}
}
