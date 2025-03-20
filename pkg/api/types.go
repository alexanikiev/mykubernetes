// File: pkg/api/types.go
package api

// Pod is a simplified representation of a Kubernetes Pod resource.
type Pod struct {
	Name   string
	Status string
}

// Service is a simplified representation of a Kubernetes Service resource.
type Service struct {
	Name string
	Type string
}
