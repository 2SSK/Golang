package main

import (
	"fmt"
	"time"
)

func factorial(n int) int {
	if n <= 1 {
		return n
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func printOutput(n int) {
	fmt.Printf("| %6d | %9d | %9d |\n", n, fibonacci(n), factorial(n))
}

func main() {
	fmt.Printf("\n| Number | Fibonacci | Factorial |\n")
	fmt.Printf("|--------|-----------|-----------|\n")
	for i := 0; i <= 10; i++ {
		go printOutput(i)
		time.Sleep(10 * time.Millisecond)
	}

	var input string
	fmt.Println("\nPress Enter to exit")
	fmt.Scanln(&input)
}
