package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	wordList := []string{"cactus","mobile","window","laptop","monitor"}
	lengthOfWords := len(wordList)
	randomNumber := generateRandomNumber(lengthOfWords)
	fmt.Println(randomNumber)	
}
func generateRandomNumber(lenWords int)int{
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	return rand.Intn(4)
}