#!/bin/bash

# File: run-golang.sh

echo "== Checking Environment for Golang =="

echo "== Preparing Go Modules =="
go mod tidy 

echo "== Building Golang Application =="
go build -o http-client || { echo "Build failed"; exit 1; }

# Run the application
echo "== Running Golang Application =="
./http-client
