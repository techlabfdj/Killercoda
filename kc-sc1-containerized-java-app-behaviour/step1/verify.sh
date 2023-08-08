#!/bin/bash
app_jar="~/killercoda-artifacts/sc1/java-mem-block-reserver-j17-latest.jar"
# Vérifier si le fichier readme.md existe
if [[ -f "$app_jar" ]]; then
  echo "Le fichier $app_jar existe."
else
  echo "Le fichier $app_jar n'existe pas."
  exit 1  # Code de sortie 1 pour indiquer une erreur
fi
