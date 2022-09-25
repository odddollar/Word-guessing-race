package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// words are 5 letter nouns, adjectives or verbs

type user struct {
	name  string
	score int
}

var users []user

func main() {
	rand.Seed(time.Now().UnixNano())

	// setup gin server
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	// generate random word list and convert to JSON
	t, _ := json.Marshal(generateWordList(readWordFile("static/words.txt")))
	wordList = string(t)
	fmt.Println(wordList)

	// host initial home page for setting username
	router.GET("/", home)

	// append new username to list of users and set initial score
	router.POST("/", homePostback)

	// handle redirection to main game page
	router.GET("/game", game)

	// run server
	router.Run("localhost:8080")
}
