#!/bin/bash

fichier="~/gitlab-runner/config/config.toml"

# Check if the specified file exists
if [ -f "$fichier" ]; then
    echo "Le fichier $fichier existe."
    exit 0 # success, return 0
else
    echo "Le fichier $fichier n'existe pas."
    exit 1 # Failed, return 1
fi