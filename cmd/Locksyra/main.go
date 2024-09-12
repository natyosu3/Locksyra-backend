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

	engine := engine.NewEngine(gin.New())

	engine.Run(":8080")
}
