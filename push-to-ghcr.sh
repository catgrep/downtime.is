#!/bin/bash

# Color codes
GREEN='\033[0;32m'
NC='\033[0m' # No
infomsg() {
    printf "${GREEN}info: %s${NC}\n" "$1" >&2
}

# Build the image & push to GitHub Container Registry
TAG=$(git rev-parse --short HEAD)
infomsg "Building and pushing Docker image with tag: $TAG"
docker buildx build \
    --progress plain \
    --platform linux/amd64,linux/arm64 \
    --build-arg VERSION="$TAG" \
    -t ghcr.io/bosdhill/downtime.is:"$TAG" \
    -t ghcr.io/bosdhill/downtime.is:latest \
    --push .
infomsg "Published: ghcr.io/bosdhill/downtime.is:$TAG"
