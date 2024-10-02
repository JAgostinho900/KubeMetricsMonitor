package main

import (
	"errors"
	"testing"

	"k8s.io/client-go/kubernetes"
)

// Mock function to simulate node fetching
func MockFetchNodes(clientset *kubernetes.Clientset) ([]string, error) {
	nodes := []string{"node1", "node2"}
	return nodes, nil
}

// Test case for FetchNodes function
func TestFetchNodes(t *testing.T) {
	mockClientset := &kubernetes.Clientset{}
	nodes, err := MockFetchNodes(mockClientset)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedNodes := 2
	if len(nodes) != expectedNodes {
		t.Errorf("Expected %d nodes, but got %d", expectedNodes, len(nodes))
	}
}

// Mock function to simulate pod fetching
func MockFetchPods(clientset *kubernetes.Clientset) ([]string, error) {
	pods := []string{"pod1", "pod2", "pod3"}
	return pods, nil
}

// Test case for FetchPods function
func TestFetchPods(t *testing.T) {
	mockClientset := &kubernetes.Clientset{}
	pods, err := MockFetchPods(mockClientset)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedPods := 3
	if len(pods) != expectedPods {
		t.Errorf("Expected %d pods, but got %d", expectedPods, len(pods))
	}
}

// Test case for handling errors
func TestFetchNodesWithError(t *testing.T) {
	mockClientset := &kubernetes.Clientset{}

	// Simulate an error in fetching nodes
	mockFetchNodesWithError := func(clientset *kubernetes.Clientset) ([]string, error) {
		return nil, errors.New("failed to fetch nodes")
	}

	_, err := mockFetchNodesWithError(mockClientset)

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
