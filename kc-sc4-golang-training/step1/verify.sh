#!/bin/bash

# Define the paths to the directories
GO_SRC="$HOME/go/src"
GO_BIN="$HOME/go/bin"

# Check if the $HOME/go/src directory exists
if [ -d "$GO_SRC" ]; then
    echo "The directory $GO_SRC exists."
else
    echo "The directory $GO_SRC does not exist."
    exit 1 # Failure, return 1
fi

# Check if the $HOME/go/bin directory exists
if [ -d "$GO_BIN" ]; then
    echo "The directory $GO_BIN exists."
else
    echo "The directory $GO_BIN does not exist."
    exit 1 # Failure, return 1
fi


# Check if Go is installed
if command -v go &> /dev/null; then
    # Get the Go version
    go_version=$(go version)

    # Check if the Go version matches "go1.21" or higher
    if [[ $go_version == *"go1.21"* ]]; then
        echo "Go 1.21 or higher is installed: $go_version"
    else
        echo "Go is installed, but it's not version 1.21 or higher: $go_version"
        exit 1 # Failure, return 1
    fi
else
    echo "Go is not installed on this system."
    exit 1 # Failure, return 1
fi

# Exit with the return code indicated by status
exit 0
