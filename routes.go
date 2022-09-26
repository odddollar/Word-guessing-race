package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// holds data from score postback
type scoreJSON struct {
	Username string `json:"name"`
}

// holds user data for internal management and posting to /score endpoint
type user struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// internal "database" of users and their scores
var users []user

// update score for user
func updateScore(username string) {
	for i := 0; i < len(users); i++ {
		if users[i].Name == username {
			users[i].Score++
			break
		}
	}
}

// get user's score from internal database
func getScore(username string) int {
	for i := 0; i < len(users); i++ {
		if users[i].Name == username {
			return users[i].Score
		}
	}
	return 0
}

// return home.html template
func home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", gin.H{})
}

// get username from form postback and add to internal database
// then redirect to relevant game page with username as query-string
func homePostback(ctx *gin.Context) {
	username := ctx.PostForm("username")
	users = append(users, user{Name: username, Score: 0})

	fmt.Println(users)

	ctx.Redirect(http.StatusSeeOther, "/game/"+username)
}

// take username query-string and word list and render game.html template
func game(ctx *gin.Context) {
	username := ctx.Param("username")

	url := ctx.Request.Host

	ctx.HTML(http.StatusOK, "game.html", gin.H{
		"WordList": string(wordList),
		"Username": username,
		"Score":    getScore(username),
		"URL":      url,
	})
}

// marshal user array to json and return as json response
func score(ctx *gin.Context) {
	body, _ := json.Marshal(users)

	ctx.JSON(http.StatusOK, string(body))
}

// bind request json to struct and update score for that user
func scorePostback(ctx *gin.Context) {
	body := scoreJSON{}
	ctx.BindJSON(&body)

	updateScore(body.Username)

	fmt.Println(users)
}
