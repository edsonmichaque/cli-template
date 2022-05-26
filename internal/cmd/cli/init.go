package kpm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "create cli-template package sources",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("init")
			return nil
		},
	}

	return cmd
}
