package main

import (
	"fmt"
	"log"
	"net/http"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	// Start the metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Use in-cluster configuration instead of kubeconfig
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error building in-cluster config: %v", err)
	}

	// Create Kubernetes client using in-cluster config
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

	// Start the HTTP server
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
