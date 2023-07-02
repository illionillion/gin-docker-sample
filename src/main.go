package main

import	"github.com/gin-gonic/gin"

func main() {
	e := gin.Default()
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
		})
	})
	e.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
			"message": "hello",
		})
	})
	e.Run(":8000")
}