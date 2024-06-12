package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InputData struct {
	Input string `json:"input" binding:"required"`
}

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLFiles("index.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/validate", func(ctx *gin.Context) {
		// var input InputData
		// if err := ctx.ShouldBindJSON(&input); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad input"})
		// 	return
		// }
		var text string
		ctx.GetString(text)

		ctx.JSON(http.StatusOK, gin.H{"input": text})
	})

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run("localhost:8080")
}
