#!/bin/bash

app_jar="$HOME/killercoda-artifacts/sc1/java-mem-block-reserver-j17-latest.jar"

# Check if the app_jar file exists
if [[ -f "$app_jar" ]]; then
  echo "The file $app_jar exists."
else
  echo "The file $app_jar does not exist."
  exit 1  # Exit code 1 to indicate an error
fi

# Check if the Docker image exists locally
if docker images | grep -q "ghcr.io/techlabfdj/go-mem-block-reserver"; then
  echo "The Docker image ghcr.io/techlabfdj/go-mem-block-reserver exists locally."
else
  echo "The Docker image ghcr.io/techlabfdj/go-mem-block-reserver does not exist locally."
  exit 2  # Exit code 2 to indicate another error
fi
