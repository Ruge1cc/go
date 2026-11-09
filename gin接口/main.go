package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 访问 /hi/你的名字，返回一句问候
	r.GET("/hi/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{"msg": "你好，" + name})
	})

	r.Run(":8080")
}
