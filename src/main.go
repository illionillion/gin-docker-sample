package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"status": "success",
		// })
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website!!",
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