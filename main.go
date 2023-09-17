package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

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

	r.POST("/add/user", func(ctx *gin.Context) {
		var requestBody Body
		err := ctx.ShouldBindJSON(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"user.name": requestBody.Name,
			"user.age":  requestBody.Age,
		})
	})

	r.Run(":8080")
}
