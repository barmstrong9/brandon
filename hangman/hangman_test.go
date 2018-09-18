package main
import(
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestRandomNumber(t *testing.T){
	lengthOfWords := 5
	result := generateRandomNumber(lengthOfWords)
	assert.True(t, isLessThan(result, lengthOfWords),"unexpected result")
}
func isLessThan(num1, num2 int)bool{
	return num1 < num2 
		
}