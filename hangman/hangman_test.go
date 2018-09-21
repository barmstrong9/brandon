package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomNumber(t *testing.T) {
	lengthOfWords := 5
	result := generateRandomNumber(lengthOfWords)
	assert.True(t, isLessThan(result, lengthOfWords), "unexpected result")
}
func isLessThan(num1, num2 int) bool {
	return num1 < num2

}
func TestValidateLetter(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"a", true},
		{"aa", false},
		{"!", false},
		{"9", false},
	}
	for _, test := range tests {
		result := isValidLetter(test.input)
		assert.Equal(t, test.expected, result)
	}

}
func TestGuessLetter(t *testing.T) {
	var tests = []struct {
		word     string
		letter   string
		expected bool
	}{
		{"aaa", "a", true},
		{"aaa", "b", false},
		{"aaa", "A", true},
		{"Abb", "a", true},
		{"AAA", "a", true},
	}
	for _, test := range tests {
		result := guessLetter(test.word, test.letter)
		assert.Equal(t, test.expected, result, test.letter)
	}
}

func TestReplaceDash(t *testing.T) {
	var tests = []struct {
		word       string
		dashedword string
		letter     string
		expected   string
	}{
		{"aaa", "---", "a", "aaa"},
		{"aaa", "---", "A", "aaa"},
		{"aaa", "---", "a", "aaa"},
		{"aaa", "---", "b", "---"},
		{"abc", "---", "b", "-b-"},
	}

	for _, test := range tests {
		result := replaceDash(test.word, test.dashedword, test.letter)

		assert.Equal(t, test.expected, result, test.letter)
	}
}