#!/bin/bash

# Check if the file exists
if [ -e "RUNNER_TOKEN" ]; then
    # Check if the file is not empty
    if [ -s "RUNNER_TOKEN" ]; then
        echo "File RUNNER_TOKEN exists and is not empty."
        exit 0  # Return 0 for success
    else
        echo "File RUNNER_TOKEN exists but is empty."
        exit 1  # Return 1 for failure
    fi
else
    echo "File RUNNER_TOKEN does not exist."
    exit 1  # Return 1 for failure
fi
