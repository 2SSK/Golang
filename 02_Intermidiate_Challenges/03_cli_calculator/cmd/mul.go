package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var mulCmd = &cobra.Command{
	Use:   "mul [numbers...]",
	Short: "Multiply two or more numbers",

	Run: func(cmd *cobra.Command, args []string) {
		// Set the first number as the result
		result := numArr[0]
		// Multiply the rest of the numbers from the first number
		for _, v := range numArr[1:] {
			result *= v
		}

		fmt.Printf("Result: %.2f\n", result)
	},
}

func init() {
	rootCmd.AddCommand(mulCmd)
}
