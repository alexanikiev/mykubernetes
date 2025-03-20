// File: pkg/printer/printer.go
package printer

import (
	"fmt"

	"alexanikiev.dev/mykubectl/pkg/api"
)

// PrintPods prints a list of Pod resources in a tabular format.
func PrintPods(pods []api.Pod) {
	fmt.Println("NAME\tSTATUS")
	for _, pod := range pods {
		fmt.Printf("%s\t%s\n", pod.Name, pod.Status)
	}
}

// PrintServices prints a list of Service resources in a tabular format.
func PrintServices(services []api.Service) {
	fmt.Println("NAME\tTYPE")
	for _, svc := range services {
		fmt.Printf("%s\t%s\n", svc.Name, svc.Type)
	}
}
