package main

import (
	"fmt"
	"unicode"
)

// Function to check Palindrome
func isPalindrome(s string) bool {
	// Normalize the string: convert to lowercase and remove non-alphanumeric characters
	var normalized []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			normalized = append(normalized, unicode.ToLower(ch))
		}
	}

	// Use two-pointer technique to check for palindrome
	left, right := 0, len(normalized)-1
	for left < right {
		if normalized[left] != normalized[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Function to check AlphaNumeric Value
func isAlphaNumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main() {
	fmt.Println("\n-----------------------------------------")
	fmt.Println("\t    Palindrome Checker")
	fmt.Printf("-----------------------------------------\n")

	// Examples for testing
	examples := []string{
		"OpenAI",
		"1234321",
		"1021",
		"This is not a Palindrome",
		"Rator",
		"Coding is fun",
		"Was it a car or a cat i saw?",
		"Eva, can i see  bees in a cave?",
	}

	// Max length of examples
	maxSize := 0

	// Find max length of array
	for _, v := range examples {
		size := len(v)
		maxSize = max(maxSize, size)
	}

	// Iteratively prints formatted output
	for _, v := range examples {
		fmt.Printf("%-*s : %t\n", maxSize, v, isPalindrome(v))
	}
}
