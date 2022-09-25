package main

import (
	"bufio"
	"math/rand"
	"os"
)

const totalWords = 10

var wordList string

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

func generateWordList(words []string) []string {
	selected := []string{}

	for i := 0; i < totalWords; i++ {
		index := rand.Intn(len(words))
		selected = append(selected, words[index])
	}

	return selected
}
