# KubeMetricsMonitor

KubeMetricsMonitor is a Go-based microservice that scrapes and exposes Kubernetes cluster metrics via an HTTP API. It is designed for cloud-native environments and integrates seamlessly with tools like Prometheus and Grafana for real-time monitoring and visualization.

## Features

- Scrapes Kubernetes cluster metrics (nodes, pods, CPU, memory usage)
- Exposes metrics through an HTTP REST API
- Containerized using Docker for portability
- Deployed and managed on Kubernetes clusters
- Integrated with Prometheus for monitoring and Grafana for visualization
- Infrastructure as Code (IaC) provisioning with Terraform
- Continuous Integration and Continuous Deployment (CI/CD) pipelines for automated testing and deployment

## Technology Stack

- **Go**: Backend microservice
- **Kubernetes**: Orchestrating and deploying the microservice
- **Prometheus**: Metrics scraping and alerting
- **Grafana**: Visualizing metrics
- **Docker**: Containerization of the microservice
- **Terraform**: Infrastructure as Code for cloud resources
- **CI/CD**: Automated deployment pipeline using GitHub Actions (or Jenkins)

## Architecture

The microservice uses the Kubernetes API to gather real-time metrics such as CPU usage, memory utilization, pod status, and node health. It exposes this data via an HTTP endpoint, making it easy to integrate into monitoring systems like Prometheus. The service can be containerized and deployed in a Kubernetes cluster for production-ready scalability.

### Prerequisites

- Go
- Docker
- Kubernetes
- Prometheus
- Grafana

### Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/kube-metrics-monitor.git
    cd kube-metrics-monitor
    ```

2. **Build the Go microservice**:
    ```bash
    go build -o kube-metrics-monitor main.go
    ```

3. **Run the microservice** locally:
    ```bash
    ./kube-metrics-monitor
    ```

4. **Run in Docker**:
    ```bash
    docker build -t your-username/kube-metrics-monitor .
    docker run -p 8080:8080 your-username/kube-metrics-monitor
    ```

5. **Deploy to Kubernetes**:
    - Apply the Kubernetes YAML manifests for the deployment and service.

6. **Integrate with Prometheus and Grafana**:
    - Configure Prometheus to scrape the microservice API for metrics.
    - Set up Grafana dashboards to visualize the data.

## Usage

Once the service is running, it will expose metrics on an HTTP API:

- **/metrics**: Returns Kubernetes cluster metrics (e.g., CPU, memory, pod statuses).

### Example:

```bash
curl http://localhost:8080/metrics
