// File: pkg/client/fake.go
package client

import "alexanikiev.dev/mykubectl/pkg/api"

// FakeClient is a mock implementation of Client, returning dummy data instead of real API calls.
type FakeClient struct {
	// In a real client, you might store config or state (e.g., auth info, server URL).
	// Here, no fields are necessary.
}

// NewFakeClient creates a new FakeClient. In a real scenario, you might pass config or context.
func NewFakeClient() Client {
	return &FakeClient{}
}

// ListPods returns a static list of pods (mocked data).
func (c *FakeClient) ListPods() []api.Pod {
	// Simulate an API call with static data
	return []api.Pod{
		{Name: "pod-1", Status: "Running"},
		{Name: "pod-2", Status: "Pending"},
	}
}

// ListServices returns a static list of services (mocked data).
func (c *FakeClient) ListServices() []api.Service {
	return []api.Service{
		{Name: "service-1", Type: "ClusterIP"},
		{Name: "service-2", Type: "NodePort"},
	}
}
