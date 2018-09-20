package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var correctCounter int
var incorrectCounter int
var isDone bool
var isLost bool
var guessedLetter []string

func main() {
	gL := 0
	lostCounter := 10
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
		if isDone == true {
			fmt.Println(dashWord)
			fmt.Println("You Have Won!")
			fmt.Println("")
			break
		}
		if isLost == true {
			fmt.Println("You Have Lost!")
			fmt.Println("The correct word was", wordToGuess)
			fmt.Println("")
			break
		}
		fmt.Printf("Guess a letter:")
		fmt.Println("")

		fmt.Println(dashWord)

		fmt.Scanln(&input)

		isValid := isValidLetter(input)
		if isValid {
			guessedLetter = append(guessedLetter, input)
			fmt.Println("you have guessed", guessedLetter[gL])
			if guessLetter(wordToGuess, input) {
				drawHangman(incorrectCounter)
				fmt.Println("correct")
				dashWord = replaceDash(wordToGuess, dashWord, input)
			} else if guessLetter(wordToGuess, input) == false {
				fmt.Println("Incorrect")
				drawHangman(incorrectCounter)
				incorrectCounter++
				fmt.Println(incorrectCounter)
				if incorrectCounter == lostCounter {
					isLost = true
				}

			}
		} else {
			fmt.Println("Not a valid input")
		}
		fmt.Println(guessedLetter)
		gL++
		input = ""
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
	for i, c := range wordToGuess {
		if string(c) == lowerCaseCharacter {

			dashWord = dashWord[:i] + string(lowerCaseCharacter) + dashWord[i+1:]
			correctCounter++
			fmt.Println(correctCounter)
			if correctCounter == int(len(wordToGuess)) {
				isDone = true
			}
		}
	}
	return dashWord
}
func drawHangman(incorrectCounter int) {
	switch incorrectCounter {
	case 0:
		fmt.Println("________")
	case 1:
		fmt.Println("+")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Printf("+")
		fmt.Printf("________")
		fmt.Println("")

	case 2:
		fmt.Println("+-----+")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Printf("+")
		fmt.Printf("________")
		fmt.Println("")

	case 3:
		fmt.Println("+-----+")
		fmt.Println("|     |")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Printf("+")
		fmt.Printf("________")
		fmt.Println("")

	case 4:
		fmt.Println("+-----+")
		fmt.Println("|     |")
		fmt.Println("|     O")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Printf("+")
		fmt.Printf("________")
		fmt.Println("")

	case 5:
		fmt.Println("+-----+")
		fmt.Println("|     |")
		fmt.Println("|     O")
		fmt.Println("|     |")
		fmt.Println("|     |")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Printf("+")
		fmt.Printf("________")
		fmt.Println("")

	case 6:
		fmt.Println("+-----+")
		fmt.Println("|     |")
		fmt.Println("|     O")
		fmt.Println("|    /|")
		fmt.Println("|     |")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Printf("+")
		fmt.Printf("________")
		fmt.Println("")

	case 7:
		fmt.Println("+-----+")
		fmt.Println("|     |")
		fmt.Println("|     O")
		fmt.Println("|    /|\\")
		fmt.Println("|     |")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Printf("+")
		fmt.Printf("________")
		fmt.Println("")

	case 8:
		fmt.Println("+-----+")
		fmt.Println("|     |")
		fmt.Println("|     O")
		fmt.Println("|    /|\\")
		fmt.Println("|     |")
		fmt.Println("|    /")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Printf("+")
		fmt.Printf("________")
		fmt.Println("")
	case 9:
		fmt.Println("+-----+")
		fmt.Println("|     |")
		fmt.Println("|     O")
		fmt.Println("|    /|\\")
		fmt.Println("|     |")
		fmt.Println("|    / \\")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Printf("+")
		fmt.Printf("________")
		fmt.Println("")
	}
}
