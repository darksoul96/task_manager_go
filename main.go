package main

import (
	"fmt"
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
		var text string
		var responseText string
		text = ctx.PostForm("command")
		fmt.Println(text)
		if text == "help" {
			responseText = HelpMessage()
		}
		ctx.Data(http.StatusOK, "text/html, charset=utf-8", []byte("<li>"+responseText+"</li>"))
	})

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run("localhost:8080")
}
