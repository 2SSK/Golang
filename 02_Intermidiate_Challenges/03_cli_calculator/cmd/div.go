package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var divCmd = &cobra.Command{
	Use:   "div [numbers...]",
	Short: "Divide two or more numbers",

	Run: func(cmd *cobra.Command, args []string) {
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
