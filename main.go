package main

import (
	"fmt"
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	var kubeconfig string

	// Locate kubeconfig file (default location is ~/.kube/config)
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	// Build configuration from kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// Create Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	fmt.Println("Application started successfully")
	fmt.Println("Successfully created Kubernetes client")

	// Fetch and print nodes
	nodes, err := FetchNodes(clientset)
	if err != nil {
		log.Fatalf("Error fetching nodes: %s", err.Error())
	}
	PrintNodes(nodes)

	// Fetch and print pods from the default namespace
	pods, err := FetchPods(clientset, "default")
	if err != nil {
		log.Fatalf("Error fetching pods: %s", err.Error())
	}
	PrintPods(pods)
}
