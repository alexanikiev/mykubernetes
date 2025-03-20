// File: pkg/cli/root.go
package cli

import (
	"github.com/spf13/cobra"
)

// NewKubectlCommand creates the root command for the CLI.
// It defines the overall CLI name and persistent flags, and attaches subcommands.
func NewKubectlCommand() *cobra.Command {
	// Define the root command
	rootCmd := &cobra.Command{
		Use:   "mykubectl",
		Short: "A simplified kubectl CLI",
		Long: `A simplified version of kubectl for illustration purposes. 
This CLI is structured similarly to kubectl, using Cobra for command parsing 
and a factory pattern for managing CLI context.`,
		// No Run field here because rootCmd itself doesn't execute an action
		// unless a subcommand is provided.
	}

	// Persistent (global) flag example: --kubeconfig for specifying config file
	rootCmd.PersistentFlags().StringP("kubeconfig", "k", "", "Path to the kubeconfig file (optional)")

	// Add subcommands to the root. In a larger project, these would be in their own files.
	rootCmd.AddCommand(NewGetCommand())     // `mykubectl get ...`
	rootCmd.AddCommand(NewVersionCommand()) // `mykubectl version`

	return rootCmd
}
