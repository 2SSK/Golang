package main

import "fmt"

// Function to reverse  an array or slice
func reverseSlice(s []int) {
	size := len(s)

	for i := 0; i < size/2; i++ {
		s[i], s[size-i-1] = s[size-i-1], s[i]
	}
}

// Function to print array in a readable format
func printArr(s []int) {
	for i, v := range s {
		if i == 0 {
			fmt.Print("| ", v)
		} else if i == len(s)-1 {
			fmt.Print(" | ", v, " |")
		} else {
			fmt.Print(" | ", v)
		}
	}
}

func print(s []int) {
	// Print original array
	fmt.Print("\nOriginal array : ")
	printArr(s)

	// Call the function
	reverseSlice(s)

	// Print reversed array
	fmt.Print("\nReversed array : ")
	printArr(s)
}

func main() {
	// Declare an array/slice
	array := []int{1, 2, 3, 4, 5}
	print(array)

	fmt.Println()

	mySlice := array[3:5]
	print(mySlice)

	// Leave a blank line at end for better visual
	fmt.Println("")
}
