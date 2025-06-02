package cli

import (
	"feed/internal/app"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "data-load",
	Short: "Loading data",
	Long:  `Loading data from the main database.`,
	Run: func(cmd *cobra.Command, args []string) {
		app.Load(appCtx)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
