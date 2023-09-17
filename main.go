package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK, gin.H{
				"message": "Hello and welcome to my web server",
			})
	})

	r.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "creating of the user",
			"id":      id,
		})
	})

	r.GET("/user", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "creating of the user",
			"id":      id,
		})
	})

	r.Run(":8080")
}
