#!/bin/bash

# Script to manage Docker Compose for Agen EDC application

echo "Stopping Docker Compose services..."
docker compose down

echo "Building Docker Compose services..."
docker compose build

echo "Starting Docker Compose services..."
docker compose up -d

echo "Starting Docker Compose log..."
docker compose logs -f

echo "Services are running. Check logs with: docker compose logs -f"
