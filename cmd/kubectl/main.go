// File: cmd/kubectl/main.go
package main

import (
	"fmt"
	"os"

	"alexanikiev.dev/mykubectl/pkg/cli"
)

func main() {
	// Initialize the root command for the CLI
	rootCmd := cli.NewKubectlCommand()

	// Execute the root command, which parses arguments and runs the appropriate subcommand
	if err := rootCmd.Execute(); err != nil {
		// In case of an error, print it to stderr and exit with a non-zero status
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
