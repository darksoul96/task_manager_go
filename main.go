package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))
// 	http.HandleFunc("/", rootHandler)
// 	http.ListenAndServe(":8080", nil)
// }
//
// // Handlers
//
// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "index.html")
// }
//

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLFiles("index.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run("localhost:8080")
}
