package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// words are 5 letter nouns, adjectives or verbs

func main() {
	rand.Seed(time.Now().UnixNano())

	if l := os.Getenv("DOCKER"); l != "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// setup gin server
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	// generate random word list and convert to JSON
	t, _ := json.Marshal(generateWordList(readWordFile("static/words.txt")))
	wordList = string(t)
	fmt.Println(wordList)

	// host initial home page for setting username
	router.GET("/", home)

	// handle postback for setting username
	router.POST("/", homePostback)

	// host "how to play page"
	router.GET("/how-to-play", howToPlay)

	// handle redirection to main game page
	router.GET("/game/:username", game)

	// api endpoint for getting scoreboard data
	router.GET("/score", score)

	// handle updating scores for user
	router.POST("/score", scorePostback)

	// add static route
	router.Static("/assets", "./static")

	// add error route
	router.NoRoute(notFound)

	// run server
	router.Run(":8080")
}
