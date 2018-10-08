package main

import "github.com/gin-gonic/gin"
import "fmt"
abcd

func getting() {
	fmt.Printf("hello, world\n")
}

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080

	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	getting()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// router.GET("/someGet", getting)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
