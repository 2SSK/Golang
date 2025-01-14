package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [numbers...]",
	Short: "Add two or more numbers",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Please provide at least two numbers to add")
			return 
		}

		result := 0.0
		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Printf("Failed to parse %s as float64\n", v)
				continue
			}

			result += num
		}

		fmt.Printf("Result: %.2f\n", result)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
