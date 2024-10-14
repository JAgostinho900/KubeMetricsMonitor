# Change the value in {} to the namespace of grafana

# Port-forward Grafana to access it locally
echo "Port-forwarding Grafana to localhost:3000..."

# Get the Grafana admin password
echo "Grafana admin password: admin"

kubectl port-forward {{svc/grafana}} 3000:3000

