/*
	Calculator command line interface
	Copyright (c) 2021 - Pratik Pradhan
*/

package cmd

import (
	"github.com/spf13/cobra"
)

type RootCmd struct {
	Cmd *cobra.Command
}

func NewRootCmd() *RootCmd {
	c := &RootCmd{
		Cmd: &cobra.Command{
			Use:   "cal",
			Short: "basic calculator",
			Long:  "",
		},
	}
	return c
}
