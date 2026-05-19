package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/karledenstal/loop/pages"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		pages.IndexPage().Render(c.Request.Context(), c.Writer)
	})

	if err := router.Run(":4282"); err != nil {
		log.Fatal(err)
	}
}
