package utils

import (
	"fmt"
	"os"
	"strconv"
)

// Make sure the arguments should be greater or equal to 2
func VerifyArgs(args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide at least two numbers")
		return
	}
}

// Parse all arguments to Float64
func ParseArgs(args []string) (numArr []float64) {
	numArr = make([]float64, 0, len(args))
	for _, v := range args {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("Failed to parse \"%s\" as float64\n", v)
			os.Exit(1)
		}
		numArr = append(numArr, num)
	}

	return numArr
}
