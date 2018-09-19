package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func main() {
	var counter int
	fmt.Println("Welcome To Hangman")
	wordList := []string{"cactus", "mobile", "window", "laptop", "monitor"}
	lengthOfWords := len(wordList)
	randomNumber := generateRandomNumber(lengthOfWords)
	fmt.Println(randomNumber)
	wordToGuess := wordList[randomNumber]
	log.Println(wordToGuess)
	input := ""

	var dashWord string
	for wordLength := 0; wordLength < int(len(wordToGuess)); wordLength++ {
		dashWord += "-"
	}

	for {
		fmt.Printf("Guess a letter:")
		fmt.Println("")

		fmt.Println(dashWord)

		fmt.Scanln(&input)

		isValid := isValidLetter(input)
		if isValid {
			if guessLetter(wordToGuess, input) {
				fmt.Println("correct")
				dashWord = replaceDash(wordToGuess, dashWord, input)
			} else {
				fmt.Println("incorrect")
			}
			counter++
			fmt.Println(counter)
		} else {
			fmt.Println("Not a valid input")
		}
	}

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
func guessLetter(wordToGuess string, guessCharacter string) bool {
	lowerCaseCharacter := strings.ToLower(guessCharacter)
	lowerWordToGuess := strings.ToLower(wordToGuess)

	return strings.Contains(lowerWordToGuess, lowerCaseCharacter)
}
func replaceDash(wordToGuess string, dashWord string, guessCharacter string) string {
	lowerCaseCharacter := strings.ToLower(guessCharacter)
	for i, c := range wordToGuess{
		if string(c) == lowerCaseCharacter{
			dashWord = dashWord[:i] + string(lowerCaseCharacter) + dashWord[i+1:]
		}
	}

	return dashWord

}
