package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	splitMsg := strings.Split(msg, " ")
	for _, word := range splitMsg {
		var newWord []string
		for i := range word {
			newWord = append(newWord, strings.Repeat(string(word[i]), i+1))
		}
		print(strings.Join(newWord, ""))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
