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
	r.Use(gin.Logger(), gin.Recovery()) // 加装中间件
	r.GET("/test_middleware", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		// panic("test panic") // 手动触发panic来看gin.recovery的作用
		c.String(200, "%s", name)
	})
	r.Run()
}
