#!/bin/bash

# Install Helm if not already installed
if ! command -v helm &> /dev/null
then
    echo "Helm not found, installing Helm..."
    curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
else
    echo "Helm is already installed"
fi

# Add the Prometheus community Helm repo
echo "Adding Prometheus community Helm repo..."
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# Install Prometheus using Helm
echo "Installing Prometheus..."
helm install prometheus prometheus-community/prometheus --namespace default

# Wait for Prometheus pods to be ready
echo "Waiting for Prometheus to be ready..."
kubectl wait --for=condition=ready pod -l app=prometheus-server --namespace default --timeout=300s

# Port-forward Prometheus to access it locally
echo "Port-forwarding Prometheus to localhost:9090..."
kubectl port-forward svc/prometheus-server 9090:9090 &

# Add the Grafana Helm repo
echo "Adding Grafana Helm repo..."
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

# Install Grafana using Helm
echo "Installing Grafana..."
helm install grafana grafana/grafana --namespace default --set adminPassword=admin

# Wait for Grafana pods to be ready
echo "Waiting for Grafana to be ready..."
kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=grafana --namespace default --timeout=300s

# Port-forward Grafana to access it locally
echo "Port-forwarding Grafana to localhost:3000..."
kubectl port-forward svc/grafana 3000:3000 &

# Get the Grafana admin password
echo "Grafana admin password: admin"
echo "Access Grafana at http://localhost:3000"

# Final instruction message
echo "Prometheus is running at http://localhost:9090"
echo "Grafana is running at http://localhost:3000"
echo "You can log in to Grafana with username 'admin' and the password 'admin'"
