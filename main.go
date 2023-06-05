package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7878"
	}
	formatedPort := ":" + port
	log.Printf("formatedPort -> " + formatedPort)

	var mode = os.Getenv("GIN_MODE")
	log.Printf("mode -> " + mode)
	if mode == "" {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(route.Run(formatedPort)) // listen and serve on 0.0.0.0:8080
}
