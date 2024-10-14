# Change the value in {} to the namespace of prometheus-server

# Port-forward Prometheus to access it locally
echo "Port-forwarding Prometheus to localhost:9090..."
kubectl port-forward {{svc/prometheus-server}} 9090:9090