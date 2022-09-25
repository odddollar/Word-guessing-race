package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", gin.H{})
}

func homePostback(ctx *gin.Context) {
	username := ctx.PostForm("username")
	users = append(users, user{name: username, score: 0})

	fmt.Println(users)

	ctx.Redirect(http.StatusSeeOther, "/game?username="+username)
}

func game(ctx *gin.Context) {
	username := ctx.Query("username")

	ctx.HTML(http.StatusOK, "game.html", gin.H{
		"WordList": string(wordList),
		"Username": username,
	})
}
