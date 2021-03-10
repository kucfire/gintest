package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Use(IPAuthMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})
	r.Run(":8080")
}

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		iplist := []string{
			"127.0.0.2",
		}
		clientIP := c.ClientIP()
		flag := false
		for _, host := range iplist {
			if clientIP == host {
				flag = true
				break
			}
		}

		if !flag {
			c.String(401, "%s, not in white list", clientIP)
			c.Abort()
		}
	}
}
