package main

import "cal/cmd"

func main() {

	rootCmd := cmd.InitialiseRootCmd()

	rootCmd.Cmd.Execute()
}
