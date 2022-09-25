package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type scoreJSON struct {
	Username string `json:"name"`
}

type user struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

var users []user

func updateScore(username string) {
	// update score for user
	for i := 0; i < len(users); i++ {
		if users[i].Name == username {
			users[i].Score++
			break
		}
	}
}

func home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", gin.H{})
}

func homePostback(ctx *gin.Context) {
	username := ctx.PostForm("username")
	users = append(users, user{Name: username, Score: 0})

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

func score(ctx *gin.Context) {
	body, _ := json.Marshal(users)

	ctx.JSON(http.StatusOK, string(body))
}

func scorePostback(ctx *gin.Context) {
	body := scoreJSON{}
	ctx.BindJSON(&body)

	updateScore(body.Username)

	fmt.Println(users)
}
