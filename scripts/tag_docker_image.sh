#!/bin/bash

# Define the Docker repository as a variable
DOCKER_REPOSITORY="{{ your repository }}"
DOCKER_TAG="latest"

# Login into Docker Hub
echo "Logging in to Docker Hub..."
docker login -u <your_dockerhub_username> -p <your_dockerhub_password>

# Tag the Docker image
echo "Tagging the Docker image..."
docker tag kube-metrics-monitor $DOCKER_REPOSITORY:$DOCKER_TAG

# Push the image to Docker Hub
echo "Pushing the image to Docker Hub..."
docker push $DOCKER_REPOSITORY:$DOCKER_TAG

echo "Docker image successfully pushed to Docker Hub."
