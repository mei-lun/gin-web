package main

import (
	"exchangeapp/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	// fmt.Printf("App Name: %s, App Port: %s\n", config.AppConfig.App.Name, config.AppConfig.App.Port)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + config.AppConfig.App.Port)
}
