package command

import "github.com/spf13/cobra"

// TODO подумать над организацией кода
var syncCmd = &cobra.Command{
	Use:   "sync-data",
	Short: "data synchronization",
	Long:  "data synchronization for feed generation",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
