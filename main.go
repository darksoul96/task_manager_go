package main

import (
	"net/http"
	"strings"

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
		commandWords := strings.Fields(command)

		switch commandWords[0] {
		case "help":
			responseText = HelpMessage()

		case "add":
			if len(commandWords) > 1 {
				AddTask(strings.Join(commandWords[1:], " "))
				responseText = "\nTask added"
			} else {
				responseText = "\nPlease add a task description"
			}

		case "list":

		default:
			responseText = "\nTask not found"
		}
		ctx.Data(http.StatusOK, "command/html, charset=utf-8", []byte("<li>"+command+responseText+"</li>"))
	})

	r.Run("localhost:8080")
}
