// File: pkg/factory/factory.go
package factory

import (
	// No external dependencies for simplicity
	"alexanikiev.dev/mykubectl/pkg/client"
)

// Factory holds configuration needed to create clients or other context for commands.
type Factory struct {
	KubeconfigPath string
}

// NewFactory creates a new Factory. If kubeconfigPath is empty, it applies a default.
func NewFactory(kubeconfigPath string) *Factory {
	if kubeconfigPath == "" {
		// Use default kubeconfig path (in a real kubectl, this might check $KUBECONFIG env or ~/.kube/config)
		kubeconfigPath = "~/.kube/config"
	}
	return &Factory{KubeconfigPath: kubeconfigPath}
}

// Client returns a client interface for interacting with the API.
// In a real scenario, this might initialize a REST client based on the kubeconfig.
func (f *Factory) Client() client.Client {
	// For the illustration, we return a fake client regardless of config, but we could use f.KubeconfigPath if needed.
	return client.NewFakeClient()
}
