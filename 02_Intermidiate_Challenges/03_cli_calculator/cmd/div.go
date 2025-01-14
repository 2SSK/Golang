package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var divCmd = &cobra.Command{
	Use:   "div [numbers...]",
	Short: "Divide two or more numbers",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Please provide at least two numbers to divide")
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

		result := numArr[0]
		for _, v := range numArr[1:] {
			if v == 0 {
				fmt.Println("Cannot divide by zero")
				return
			}
			result /= v
		}

		fmt.Printf("Result: %.2f\n", result)
	},
}

func init() {
	rootCmd.AddCommand(divCmd)
}
