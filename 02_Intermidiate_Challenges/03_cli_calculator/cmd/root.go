package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// To store parsed arguments
var numArr []float64

// Make sure the arguments should be greater or equal to 2
func verifyArgs(args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide at least two numbers")
		os.Exit(1)
	}
}

// Parse all arguments to Float64
func parseArgs(args []string) {
	numArr = make([]float64, 0, len(args))
	for _, v := range args {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("Failed to parse %s as float64\n", v)
		}
		numArr = append(numArr, num)
	}
}

var rootCmd = &cobra.Command{
	Use:   "calcli",
	Short: "A simple cli calculator",

	// Apply function to all subcommands
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		verifyArgs(args)
		parseArgs(args)
	},

	// Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}

