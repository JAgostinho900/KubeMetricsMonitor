#!/bin/bash

# Login into Docker Hub
echo "Logging in to Docker Hub..."
docker login

# Tag the Docker image
echo "Tagging the Docker image..."
docker tag kube-metrics-monitor jagostinho900/kube-metrics-monitor:latest

# Push the image to Docker Hub
echo "Pushing the image to Docker Hub..."
docker push jagostinho900/kube-metrics-monitor:latest

echo "Docker image successfully pushed to Docker Hub."
