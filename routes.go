package main

import (
	"encoding/json"
	"fmt"

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
	return -1
}

// return home.html template
func home(ctx *gin.Context) {
	ctx.HTML(200, "home.html", gin.H{})
}

// get username from form postback and add to internal database
// then redirect to relevant game page with username as query-string
func homePostback(ctx *gin.Context) {
	username := ctx.PostForm("username")
	users = append(users, user{Name: username, Score: 0})

	fmt.Println(users)

	ctx.Redirect(303, "/game/"+username)
}

// take username query-string and word list and render game.html template
func game(ctx *gin.Context) {
	username := ctx.Param("username")

	url := ctx.Request.Host

	// check that user is valid and return 404 if not
	s := getScore(username)
	if s == -1 {
		notFound(ctx)
	} else {
		ctx.HTML(200, "game.html", gin.H{
			"WordList": string(wordList),
			"Username": username,
			"Score":    s,
			"URL":      url,
		})
	}
}

// marshal user array to json and return as json response
func score(ctx *gin.Context) {
	body, _ := json.Marshal(users)

	ctx.JSON(200, string(body))
}

// bind request json to struct and update score for that user
func scorePostback(ctx *gin.Context) {
	body := scoreJSON{}
	ctx.BindJSON(&body)

	updateScore(body.Username)

	fmt.Println(users)
}

func notFound(ctx *gin.Context) {
	ctx.HTML(404, "error.html", gin.H{
		"Error":   "404",
		"Message": fmt.Sprint("\"" + ctx.Request.Host + ctx.Request.URL.Path + "\" not found"),
	})
}
