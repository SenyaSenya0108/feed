package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add numbers",
	Long:  `Adds two numbers and prints the result.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding numbers...")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
