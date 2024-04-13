package router

import (
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	router.Run(":8000")
}
