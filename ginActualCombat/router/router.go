package main

import "github.com/gin-gonic/gin"

func main() {
	c := gin.Default()
	c.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	c.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})
	c.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})
	c.Any("/any", func(c *gin.Context) {
		c.String(200, "Any")
	})
	c.Run()
}
