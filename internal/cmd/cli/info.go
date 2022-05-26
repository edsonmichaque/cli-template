package kpm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "installs cli-template package",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("info")
			return nil
		},
	}

	return cmd
}
