#!/bin/bash

# Clean previous builds
echo "Cleaning previous builds..."
rm -rf build

# Create a new build directory
mkdir build

# Build the project for ARM64
echo "Building the project for ARM64..."
GOARCH=arm64 GOOS=sylixos go build -o build/sylix-go-webserver-arm64

# Check if the ARM64 build was successful
if [ $? -eq 0 ]; then
    echo "ARM64 build successful!"
else
    echo "ARM64 build failed!"
    exit 1
fi

# Build the project for AMD64
echo "Building the project for AMD64..."
GOARCH=amd64 GOOS=sylixos go build -o build/sylix-go-webserver-amd64

# Check if the AMD64 build was successful
if [ $? -eq 0 ]; then
    echo "AMD64 build successful!"
else
    echo "AMD64 build failed!"
    exit 1
fi