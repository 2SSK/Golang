package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub [numbers...]",
	Short: "Subtract two or more numbers",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Please provide at least two numbers to subtract")
		}

		// Parse the arguments
		numArr := make([]float64, 0, len(args))
		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Printf("Failed to parse %s as float64\n", v)
			}
			numArr = append(numArr, num)
		}

		// Set the first number as the result
		result := numArr[0]

		// Subtract the rest of the numbers from the first number
		for _, v := range numArr[1:] {
			result -= v
		}

		fmt.Printf("Result: %.2f\n", result)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
