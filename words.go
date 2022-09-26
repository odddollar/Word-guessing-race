package main

import (
	"bufio"
	"math/rand"
	"os"
)

const totalWords = 10

var wordList string

// return array of strings containing lines of file
func readWordFile(f string) []string {
	words := []string{}

	file, _ := os.Open(f)
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		words = append(words, s.Text())
	}

	return words
}

// select random elements from input string array
func generateWordList(words []string) []string {
	selected := []string{}

	for i := 0; i < totalWords; i++ {
		index := rand.Intn(len(words))
		selected = append(selected, words[index])
	}

	return selected
}
