#!/bin/bash

readme_file="readme.md"

# Vérifier si le fichier readme.md existe
if [[ -f "$readme_file" ]]; then
  echo "Le fichier $readme_file existe."

  # Vérifier si le fichier contient au moins une occurrence de "PetClinic"
  if grep -q "PetClinic" "$readme_file"; then
    echo "Le fichier $readme_file contient au moins une occurrence de 'PetClinic'."
    exit 0  # Code de sortie 0 pour indiquer que tout est OK
  else
    echo "Le fichier $readme_file ne contient pas 'PetClinic'."
    exit 1  # Code de sortie 1 pour indiquer une erreur
  fi
else
  echo "Le fichier $readme_file n'existe pas."
  exit 1  # Code de sortie 1 pour indiquer une erreur
fi
