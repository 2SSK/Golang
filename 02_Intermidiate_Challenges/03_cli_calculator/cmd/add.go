package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [numbers...]",
	Short: "Add two or more numbers",

	Run: func(cmd *cobra.Command, args []string) {
		result := numArr[0]
		for _, v := range numArr[1:] {
			result += v
		}

		fmt.Printf("Result: %.2f\n", result)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
