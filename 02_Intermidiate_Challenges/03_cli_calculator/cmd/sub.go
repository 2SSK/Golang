package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub [numbers...]",
	Short: "Subtract two or more numbers",

	Run: func(cmd *cobra.Command, args []string) {
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
