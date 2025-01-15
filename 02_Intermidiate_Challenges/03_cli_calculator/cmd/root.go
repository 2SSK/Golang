package cmd

import (
	"fmt"
	"os"

	"github.com/2SSK/Golang/02_Intermidiate_Challenges/03_cli_calculator/utils"
	"github.com/spf13/cobra"
)

var (
	addFlag bool
)

var numArr []float64

var rootCmd = &cobra.Command{
	Use:   "calcli",
	Short: "A simple cli calculator",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		utils.VerifyArgs(args)
		numArr = utils.ParseArgs(args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
