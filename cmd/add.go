/*
	Calculator command line interface
	Copyright (c) 2021 - Pratik Pradhan
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type AddCmd struct {
	Cmd *cobra.Command
}

func NewAddCmd() *AddCmd {
	c := &AddCmd{
		Cmd: &cobra.Command{
			Use:   "add",
			Short: "Add two numbers",
			Long:  "",
		},
	}
	c.Cmd.Run = c.Run
	return c
}

func (c *AddCmd) Run(cmd *cobra.Command, args []string) {
	num1 := 20
	num2 := 10
	result := num1 + num2
	fmt.Printf("Addition of %d and %d is: %d", num1, num2, result)
}
