# KubeMetricsMonitor

KubeMetricsMonitor is a Go-based microservice designed to scrape and expose Kubernetes cluster metrics via an HTTP API. This microservice is optimized for cloud-native environments and integrates seamlessly with monitoring tools like **Prometheus** and **Grafana** to provide real-time metrics and visualizations.

## Features

- Scrapes Kubernetes cluster metrics such as nodes, pods, CPU, and memory usage.
- Exposes these metrics through an HTTP REST API for easy consumption.
- Fully containerized with Docker for portability across environments.
- Deployed and managed within Kubernetes clusters for scalability.
- Integrated with Prometheus for monitoring and Grafana for visualization.
- Supports a robust CI pipeline using Jenkins for automated testing and building (currently running locally).

## Technology Stack

- **Go**: Backend microservice that interacts with the Kubernetes API.
- **Kubernetes**: Manages and orchestrates the deployment of the microservice.
- **Prometheus**: Scrapes metrics from the microservice and handles alerting.
- **Grafana**: Visualizes metrics with customizable dashboards.
- **Docker**: Ensures the service is easily portable across different environments.
- **Jenkins**: Automates the build pipeline (locally) and ensures continuous integration.

## Architecture

The service interacts with the Kubernetes API to gather real-time metrics, including resource usage (CPU and memory), pod statuses, and node health. These metrics are exposed via an HTTP endpoint (`/metrics`), making them easily accessible to monitoring tools such as Prometheus. The system is fully containerized, designed for seamless deployment in Kubernetes clusters, and supports horizontal scaling for production use.

## Key Features

- **Kubernetes API Integration**: The microservice fetches critical cluster information, enabling comprehensive monitoring.
- **HTTP API**: Metrics are made available through an accessible HTTP endpoint.
- **Dockerized Deployment**: The microservice is containerized, ensuring compatibility across environments.
- **Prometheus Integration**: Out-of-the-box support for Prometheus to scrape metrics.
- **Grafana Dashboards**: Visualize cluster health, resource utilization, and performance through Grafana.
- **Jenkins CI**: Local Jenkins setup provides a continuous integration pipeline to build and test the application.

## Deployment and Setup

### 1. Build the Docker Image
Use the following commands to build the Docker image of the Go application.

### 2. Deploy the Application to Kubernetes
Apply the Kubernetes deployment YAML file to deploy the application to your Kubernetes cluster.

### 3. Apply Kubernetes Permissions (ClusterRole and ClusterRoleBinding)
To allow the application to list nodes and pods, apply the necessary ClusterRole and ClusterRoleBinding YAML files.

### 4. Install and Run Prometheus and Grafana
Install Prometheus and Grafana using Helm, then port-forward to make them accessible locally.

### 5. Verifying the Application
Once the setup is complete, access Prometheus and Grafana locally and ensure they are running properly.

## Jenkins CI Pipeline (Local)

At the moment, Jenkins is set up locally to automate the continuous integration pipeline. It builds and tests the application. The pipeline works locally, but the Jenkins server will be migrated to a cloud environment later.

You will need to manually trigger the Jenkins builds from the Jenkins UI until further automation steps are implemented.

## Future Enhancements

- **Jenkins on AWS**: Transitioning the Jenkins server to run on an AWS EC2 instance for production-grade CI/CD pipelines.
- **Advanced Monitoring**: Adding custom metrics and fine-tuning Grafana dashboards for detailed performance insights.
- **GitHub Actions**: Implementing GitHub Actions to trigger Jenkins builds on pull requests for improved CI automation (planned for future phases).
