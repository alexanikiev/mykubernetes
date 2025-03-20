// File: pkg/cli/get.go
package cli

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/cobra"

	"alexanikiev.dev/mykubectl/pkg/api"
	"alexanikiev.dev/mykubectl/pkg/factory"
	"alexanikiev.dev/mykubectl/pkg/printer"
)

// NewGetCommand creates a new 'get' subcommand.
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [resource]",
		Short: "Display one or many resources",
		Args:  cobra.MinimumNArgs(1), // Require at least one argument (the resource type)
		Run: func(cmd *cobra.Command, args []string) {
			// Determine the resource(s) requested (e.g., "pods" or "pods,services")
			resourceArg := args[0]

			// Initialize factory (configure client using kubeconfig flag if provided)
			configPath, _ := cmd.Root().PersistentFlags().GetString("kubeconfig")
			f := factory.NewFactory(configPath)
			client := f.Client() // Get a client interface to use for API calls

			// Check if multiple resource types are requested (comma-separated)
			if strings.Contains(resourceArg, ",") {
				// For multiple resource types, we will fetch each concurrently.
				resources := strings.Split(resourceArg, ",")
				var wg sync.WaitGroup
				var pods []api.Pod
				var services []api.Service

				// Launch a goroutine for each resource type requested.
				for _, res := range resources {
					switch res {
					case "pods":
						wg.Add(1)
						go func() {
							defer wg.Done()
							pods = client.ListPods()
						}()
					case "services":
						wg.Add(1)
						go func() {
							defer wg.Done()
							services = client.ListServices()
						}()
					}
				}

				// Wait for all goroutines to finish API calls
				wg.Wait()

				// Print the results for each resource type gathered
				if pods != nil {
					printer.PrintPods(pods)
				}
				if services != nil {
					printer.PrintServices(services)
				}
			} else {
				// Single resource type handling (no concurrency needed here)
				switch resourceArg {
				case "pods":
					pods := client.ListPods()
					printer.PrintPods(pods)
				case "services":
					services := client.ListServices()
					printer.PrintServices(services)
				default:
					fmt.Printf("Unknown resource type: %s\n", resourceArg)
				}
			}
		},
	}
	return cmd
}
