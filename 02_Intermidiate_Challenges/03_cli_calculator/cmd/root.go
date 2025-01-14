package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Variables to store flag values
var (
	addFlag      []string
	subtractFlag []string
	multiplyFlag []string
	divideFlag   []string
)

var rootCmd = &cobra.Command{
	Use:   "calcli",
	Short: "A simple cli calculator",

	Run: func(cmd *cobra.Command, args []string) {
		// Handle add flag
		if cmd.Flags().Changed("add") || cmd.Flags().Changed("a") {
			addCmd.Run(addCmd, addFlag)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add flags that accept multiple values
	rootCmd.Flags().StringSliceVarP(&addFlag, "add", "a", []string{}, "Add numbers")
}
