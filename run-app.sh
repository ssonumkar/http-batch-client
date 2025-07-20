#!/bin/bash

# File: run-golang.sh

echo "== Checking Environment for Golang =="

# Check Go installation
if ! command -v go &> /dev/null; then
  echo "Install Go"
else
  echo "Go found: $(go version)"
fi

echo "== Preparing Go Modules =="
go mod tidy 

echo "== Building Golang Application =="
go build -o http-client || { echo "Build failed"; exit 1; }

# Run the application
echo "== Running Golang Application =="
./http-client
