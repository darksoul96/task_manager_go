package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLFiles("index.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/validate", func(ctx *gin.Context) {
		var responseText string
		command := ctx.PostForm("command")
		fmt.Println(command)
		switch command {
		case "help":
			responseText = HelpMessage()
		}
		ctx.Data(http.StatusOK, "command/html, charset=utf-8", []byte("<li>"+command+responseText+"</li>"))
	})

	r.Run("localhost:8080")
}
