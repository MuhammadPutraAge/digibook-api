package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadputraage/digibook-api/config"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to digibook API",
		})
	})
	r.Run()
}
