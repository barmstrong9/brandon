package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
	"log"
)

func main() {
	fmt.Println("Welcome To Hangman")
	wordList := []string{"cactus", "mobile", "window", "laptop", "monitor"}
	lengthOfWords := len(wordList)
	randomNumber := generateRandomNumber(lengthOfWords)
	fmt.Println(randomNumber)
	wordToGuess := wordList[randomNumber]
	log.Println(wordToGuess)
	input := ""
	for {
		fmt.Printf("Guess a letter:")

		fmt.Scanln(&input)
		isValid := isValidLetter(input)
			if isValid{
				break
			}else {
				fmt.Println("Not a valid input")
			}
		
	}
	fmt.Println(input)
}
func generateRandomNumber(lenWords int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	return rand.Intn(5)
}
func isValidLetter(input string) bool {
	if len(input) == 1 {
		isLetter, _ := regexp.MatchString("^[a-zA-Z]", input)
		if isLetter {
			return true
		}

	}
	return false

}
