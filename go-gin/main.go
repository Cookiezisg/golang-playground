package main

import (
	"go-gin/testFunc"

	"github.com/gin-gonic/gin"
)

func main() {
	println(testFunc.SampleFunc())
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")

}
