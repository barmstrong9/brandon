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
func TestValidateLetter(t *testing.T){
	var tests = []struct{
		input string
		expected bool
	}{
		{"a",true},
		{"aa", false},
		{"!", false},
		{"9", false},
	}
	for _, test := range tests{
		result := isValidLetter(test.input)
	assert.Equal(t, test.expected, result)
	}
	
}