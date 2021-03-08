package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/test_middleware", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(200, "%s", name)
	})
	r.Run()
}