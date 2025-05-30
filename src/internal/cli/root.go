package cli

import (
	"feed/internal/app"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Hugo is a very fast static site generator",
	Long:  `long`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root cobra")
	},
}

var appCtx *app.AppContext

func Execute(ctx *app.AppContext) {
	//Создаем контекст в пространстве пакета консольных команд, для DI
	appCtx = ctx
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
