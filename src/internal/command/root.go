package command

import (
	"github.com/spf13/cobra"
)

type CLI struct {
	cmd *cobra.Command
}

func NewRootCMD() *CLI {
	cmd := &cobra.Command{
		Use:   "feeds",
		Short: "An application that generates feeds for various services",
		Long:  "An application that generates feeds for various services",
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println()
		},
	}

	return &CLI{cmd: cmd}
}

func (c *CLI) Execute() error {
	c.cmd.AddCommand(syncCmd)
	if err := c.cmd.Execute(); err != nil {
		return err
	}

	return nil
}
