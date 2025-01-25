package main

import (
	"reflect"
	"testing"
)

type Input struct {
	arr []int
	l   int
	r   int
}

func TestMergeSort(t *testing.T) {

	tests := []struct {
		name     string
		input    Input
		expected []int
	}{
		{"Single element", Input{[]int{5}, 0, 0}, []int{5}},
		{"Already sorted", Input{[]int{1, 2, 3, 4, 5}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{"Reverse order", Input{[]int{5, 4, 3, 2, 1}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{"Unsorted array", Input{[]int{5, 2, 9, 1, 6}, 0, 4}, []int{1, 2, 5, 6, 9}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Copy the input to avoid modifying the original test case
			inputArr := append([]int(nil), test.input.arr...)

			// Run the mergeSort function
			mergeSort(inputArr, test.input.l, test.input.r)

			// Check if the result matches the expected value
			if !reflect.DeepEqual(inputArr, test.expected) {
				t.Errorf("For %s: expected %v, but got %v", test.name, test.expected, inputArr)
			}
		})
	}
}
