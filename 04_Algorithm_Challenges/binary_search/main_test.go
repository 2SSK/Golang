package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"Basic Test Case 1", []int{1, 3, 5, 7, 9}, 5, 2},
		{"Basic Test Case 2", []int{2, 4, 6, 8, 10}, 8, 3},
		{"First Element", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Last Element", []int{1, 2, 3, 4, 5}, 5, 4},
		{"Target Not Present", []int{1, 2, 3, 4, 5}, 6, -1},
		{"Empty Array", []int{}, 3, -1},
		{"Single Element - Found", []int{10}, 10, 0},
		{"Single Element - Not Found", []int{10}, 5, -1},
		{"Duplicate Elements", []int{1, 2, 2, 2, 3}, 2, 2}, // Any valid index: 1, 2, or 3
		{"All Elements Same", []int{7, 7, 7, 7, 7}, 7, 2},  // Any valid index: 0-4
		{"Large Array", generateLargeArray(1000000), 999999, 999998},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Run the binarySearch function
			output := binarySearch(test.arr, test.target)

			// Check if the result matches the expected value
			if output != test.expected {
				t.Errorf("For %s: expected %v, but got %v", test.name, test.expected, output)
			}
		})
	}
}

func generateLargeArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i + 1
	}
	return arr
}
