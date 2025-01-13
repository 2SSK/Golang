package main

import (
	"fmt"
	"math"
)

// INFO: Brute force approach
// func primeNo(num int) bool {
// 	if num <= 1 {
// 		return false
// 	}
//
// 	if num <= 3 {
// 		return true
// 	}
//
// 	if num%2 == 0 || num%3 == 0 {
// 		return false
// 	}
//
// 	for i := 5; i*i <= num; i += 6 {
// 		if num%i == 0 || num%i+2 == 0 {
// 			return false
// 		}
// 	}
//
// 	return true
// }

// INFO: (Sieve of Eratothenes) Optimal approach to find all primes less than num
func sieveOfEratothenes(num int) []int {
	// Create an array of boolean values to represent from numbers 0 to num
	isPrime := make([]bool, num+1)

	// Assign true to all elements of array
	for i := range isPrime {
		isPrime[i] = true
	}

	// 0 and 1 are not prime
	isPrime[0] = false
	isPrime[1] = false

	// Start from 2 till num and mark non-primes numbers
	for p := 2; p <= int(math.Sqrt(float64(num))); p++ {
		if isPrime[p] {
			// Mark multiples of p as not prime
			for multiple := p * p; multiple <= num; multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	// Collect all prime numbers into an array / slice
	primes := []int{}
	for i := 2; i <= num; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	// Take a number from user
	var num int
	fmt.Print("Enter a number : ")
	fmt.Scanf("%d", &num)

	primes := sieveOfEratothenes(num)

	count := 0
	fmt.Printf("\nAll primes less than %v are : \n\n", num)
	for _, v := range primes {
		if count >= 10 {
			fmt.Println()
			count = 0
		}
		fmt.Printf("%v ", v)
		count++
	}
	fmt.Println()
}
