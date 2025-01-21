package main

import (
	"fmt"
	"url-shortner/handler"
	"url-shortner/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore()

	// Start the web server
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server 0 Error: %v", err))
	}
}
