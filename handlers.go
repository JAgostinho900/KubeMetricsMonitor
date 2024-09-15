package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"                   // For Kubernetes core objects (NodeList, PodList, etc.)
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" // For metadata like ListOptions
	"k8s.io/client-go/kubernetes"
)

// FetchNodes fetches the nodes from the Kubernetes cluster
func FetchNodes(clientset *kubernetes.Clientset) (*corev1.NodeList, error) {
	return clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
}

// FetchPods fetches the pods from the Kubernetes cluster
func FetchPods(clientset *kubernetes.Clientset, namespace string) (*corev1.PodList, error) {
	return clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
}

// PrintNodes prints the details of the fetched nodes
func PrintNodes(nodes *corev1.NodeList) {
	fmt.Println("Listing Kubernetes Nodes:")
	for _, node := range nodes.Items {
		fmt.Printf("Node Name: %s\n", node.Name)
	}
}

// PrintPods prints the details of the fetched pods
func PrintPods(pods *corev1.PodList) {
	fmt.Println("Listing Kubernetes Pods:")
	for _, pod := range pods.Items {
		fmt.Printf("Pod Name: %s in Namespace: %s\n", pod.Name, pod.Namespace)
	}
}
