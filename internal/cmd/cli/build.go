package kpm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdBuild() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "builds a cli-template",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("build")
			return nil
		},
	}

	return cmd
}
