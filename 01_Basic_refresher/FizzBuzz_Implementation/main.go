package main

import (
	"fmt"
)

func main() {
	heading("FizzBuzz Implementation")

	// FizzBuzz implementation
	counts := FizzBuzzImplementation()

	heading("FizzBuzz Summary")
	summaryOutput("Multiple of 3 (Fizz)", counts[0])
	summaryOutput("Multiple of 5 (Buzz)", counts[1])
	summaryOutput("Multiple of both 3 & 5 (FizzBuzz)", counts[2])
	summaryOutput("Not a Multiple of 3 & 5", counts[3])
}

// Function to print heading
func heading(s string) {
	fmt.Println("\n--------------------------------------")
	fmt.Println("\t", s)
	fmt.Println("--------------------------------------")
}

func summaryOutput(s string, i int) {
	fmt.Printf("%-35s: %d\n", s, i)
}

// Function to implement FizzBuzz
func FizzBuzzImplementation() []int {
	fizzCount := 0
	buzzCount := 0
	fizzBuzzCount := 0

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fizzbuzzOutput(i, "Fizz")
			fizzCount++
		} else if i%3 != 0 && i%5 == 0 {
			fizzbuzzOutput(i, "Buzz")
			buzzCount++
		} else if i%3 == 0 && i%5 == 0 {
			fizzbuzzOutput(i, "FizzBuzz")
			fizzBuzzCount++
		} else {
			fizzbuzzOutput(i, "!..")
		}
	}

	notFizzBuzzCount := 100 - (fizzCount + buzzCount + fizzBuzzCount)
	return []int{fizzCount, buzzCount, fizzBuzzCount, notFizzBuzzCount}
}

// Function to give formatted the output
func fizzbuzzOutput(i int, s string) {
	fmt.Printf("\t%-4d: %s\n", i, s)
}
