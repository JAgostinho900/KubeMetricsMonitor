package main

import (
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	// Create a Kubernetes client
	clientset, err := getKubernetesClient()
	if err != nil {
		panic(err.Error())
	}

	// For now, let's just print out the clientset to ensure the connection is working
	fmt.Println("Successfully created Kubernetes client:", clientset)
}

// getKubernetesClient initializes a client for interacting with the Kubernetes API
func getKubernetesClient() (*kubernetes.Clientset, error) {
	var kubeconfig string

	// Check if we're running inside a Kubernetes cluster or outside of it
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	// Build configuration from within the cluster
	config, err := rest.InClusterConfig()
	if err != nil {
		// If we're outside the cluster, use the local kubeconfig file
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}

	// Error handling
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
