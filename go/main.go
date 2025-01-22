package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:id", func(c *gin.Context) {

		c.Redirect(301, "https://www.google.com")

		// c.JSON(200, gin.H{
		// 	"message":    c.ClientIP(),
		// 	"contentype": c.ContentType(),
		// 	"query":      c.Query("id"),
		// 	"color":      c.Query("color"),
		// 	"path":       c.FullPath(),
		// 	"id":         c.Param("id"),
		// })
	})
	r.Run()
}
