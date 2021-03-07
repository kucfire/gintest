package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"20060102"`
	Age      int       `form:"age"`
}

func main() {
	r := gin.Default()

	r.GET("/testing", Testing)
	r.POST("/testing", Testing)
	r.Run()
}

func Testing(c *gin.Context) {
	var person Person
	// 根据请求的content-type来做不同binding操作
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(200, "person binding error: %v", err)
	}
}
