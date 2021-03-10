package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello")
	})

	autotls.Run(r, "www.itpp.tk")
}
