package main

import (
	"Locksyra/pkg/db"
	"Locksyra/pkg/engine"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, Locksyra!")
	db.Connect()
	engine := engine.NewEngine()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Locksyra!",
		})
	})

	engine.Run(":8080")
}
