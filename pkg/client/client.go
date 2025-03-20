// File: pkg/client/client.go
package client

import "alexanikiev.dev/mykubectl/pkg/api"

// Client defines methods for interacting with the Kubernetes API (simplified for our use-case).
// This allows multiple implementations (real or fake).
type Client interface {
	ListPods() []api.Pod
	ListServices() []api.Service
	// Additional methods (e.g., Get, Create) could be defined here.
}
