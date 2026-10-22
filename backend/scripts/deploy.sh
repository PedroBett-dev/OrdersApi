#!/bin/bash

date=$(date '+%Y-%m-%d %H:%M:%S')

# Roda a partir da raiz do repositório, de qualquer diretório
cd "$(git rev-parse --show-toplevel)" || exit 1

git add .
git reset -q -- '.env' '*/.env' '*.log' 'logs' '*/logs' '.vscode' 2>/dev/null
git commit -m "chore: update from ${date}"
git push
