#!/bin/bash

# Check if Go is installed
if command -v go &> /dev/null; then
    # Get the Go version
    go_version=$(go version)

    # Check if the Go version matches "go1.21" or higher
    if [[ $go_version == *"go1.21"* ]]; then
        echo "Go 1.21 or higher is installed: $go_version"
        exit 0 # Success, return 0
    else
        echo "Go is installed, but it's not version 1.21 or higher: $go_version"
        exit 1 # Failure, return 1
    fi
else
    echo "Go is not installed on this system."
    exit 1 # Failure, return 1
fi
