package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "views/home.html", gin.H{})
}

func homePostback(ctx *gin.Context) {
	users = append(users, user{name: ctx.PostForm("username"), score: 0})

	fmt.Println(users)

	ctx.Redirect(http.StatusSeeOther, "/game")
}

func game(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "views/game.html", gin.H{
		"WordList": string(wordList),
	})
}
