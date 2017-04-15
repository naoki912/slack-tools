package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(newVersionCmd())
}

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the slack-tools version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.0.0")
		},
	}

	return cmd
}
