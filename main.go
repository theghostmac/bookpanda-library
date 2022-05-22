package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// create Gin router
	var router *gin.Engine
	router = gin.Default()
	router.Static("/", "./static-files")
	err := router.Run()
	if err != nil {
		return
	}
}
