#!/bin/bash

FILE=~/go/src/hello-world/hello.go

# Check if the file exists
if [[ ! -e "$FILE" ]]; then
    echo "File $FILE does not exist."
    exit 1
fi

# Check if the file is not empty
if [[ ! -s "$FILE" ]]; then
    echo "File $FILE is empty."
    exit 1
fi

# If file exists and is not empty, return 0
echo "File $FILE exists and is not empty."
exit 0
