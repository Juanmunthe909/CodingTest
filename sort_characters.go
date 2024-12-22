package main

import (
	"fmt"
	"strings"
)

func sortCharacters() {
	var input string
	fmt.Print("Input one line of words (S): ")
	fmt.Scanln(&input)

	vowels := "aeiou"
	vowelChars := ""
	consonantChars := ""

	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
	for _, char := range input {
		if strings.ContainsRune(vowels, char) {
			vowelChars += string(char)
		} else {
			consonantChars += string(char)
		}
	}

	fmt.Printf("Vowel Characters: %s\n", vowelChars)
	fmt.Printf("Consonant Characters: %s\n", consonantChars)
}

func main() {
	sortCharacters()
}
