#!/bin/bash

# Vérifier si l'application est accessible sur le port 8080 en local
response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080)

if [[ $response -eq 200 ]]; then
  echo "L'application est accessible sur le port 8080 en local."
  exit 0  # Code de sortie 0 pour indiquer que tout est OK
else
  echo "L'application n'est pas accessible sur le port 8080 en local."
  exit 1  # Code de sortie 1 pour indiquer une erreur
fi
