/*
	Calculator command line interface
	Copyright (c) 2021 - Pratik Pradhan
*/
package cmd

func InitialiseRootCmd() *RootCmd {
	rootCmd := NewRootCmd()
	addCmd := NewAddCmd()
	rootCmd.Cmd.AddCommand(addCmd.Cmd)
	return rootCmd
}
