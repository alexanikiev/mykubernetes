// File: pkg/cli/version.go
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewVersionCommand returns a Cobra command that prints the version of the CLI.
func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of mykubectl",
		Run: func(cmd *cobra.Command, args []string) {
			// In a real application, this might pull version info from build flags.
			fmt.Println("mykubectl version 0.1.0")
		},
	}
}
