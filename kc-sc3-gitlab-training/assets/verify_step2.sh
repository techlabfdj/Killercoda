#!/bin/bash

# Check if the environment variable RUNNER_TOKEN is defined and not empty
if [ -n "$RUNNER_TOKEN" ]; then
    echo "The environment variable RUNNER_TOKEN is defined and not empty."
    exit 0 # Success, return 0
else
    echo "The environment variable RUNNER_TOKEN is not defined or is empty."
    exit 1 # Failure, return 1
fi
