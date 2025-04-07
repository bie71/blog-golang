package main

import (
	"fmt"
	"os"

	"blog/cmd"
)

func main() {
	// Initialize the root command
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
