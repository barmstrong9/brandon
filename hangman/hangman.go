package main

import (
	"fmt"
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
	fmt.Println("\nWelcome To Hangman")
	fmt.Println("")
	wordList := []string{"cactus", "mobile", "window", "laptop", "monitor"}
	lengthOfWords := len(wordList)
	randomNumber := generateRandomNumber(lengthOfWords)
	wordToGuess := wordList[randomNumber]
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
		fmt.Printf("Please Guess a Letter:")
		fmt.Println("")

		fmt.Println(dashWord)

		fmt.Scanln(&input)

		isValid := isValidLetter(input)
		lowerInput := strings.ToLower(input)
		for _, letter := range guessedLetter {
			if lowerInput == letter {
				isValid = false
			}
		}
		if isValid {
			guessedLetter = append(guessedLetter, lowerInput)
			if guessLetter(wordToGuess, input) {
				fmt.Println("\ncorrect")
				fmt.Println("")
				drawHangman(incorrectCounter)
				dashWord = replaceDash(wordToGuess, dashWord, input)
				gL++
			} else if guessLetter(wordToGuess, input) == false {
				incorrectCounter++
				fmt.Println("\nIncorrect")
				fmt.Println("")
				drawHangman(incorrectCounter)
				gL++
				if incorrectCounter == lostCounter {
					isLost = true
				}

			}
		} else {
			fmt.Println("Thats Is Not A Valid Input Or You Have Already Guessed That Letter, Try Again")
			fmt.Println("")
		}
		fmt.Println("So far you have guessed:")
		fmt.Println(strings.Join(guessedLetter, ", "))
		fmt.Println("")
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
			if correctCounter == int(len(wordToGuess)) {
				isDone = true
			}
		}
	}
	return dashWord
}
func drawHangman(incorrectCounter int) {
	switch incorrectCounter {
	case 1:
		fmt.Println("________")
	case 2:
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

	case 3:
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

	case 4:
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

	case 5:
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

	case 6:
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

	case 7:
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

	case 8:
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

	case 9:
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
	case 10:
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