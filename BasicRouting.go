package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello World")
	})

	router.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(200, "UserID:"+id)
	})

	router.GET("/search", func(ctx *gin.Context) {
		query := ctx.DefaultQuery("q", "default-value")
		ctx.String(200, "Search query"+query)
	})

	router.Run(":8080")
}
