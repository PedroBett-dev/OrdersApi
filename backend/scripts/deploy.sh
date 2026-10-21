#!/bin/bash

date=$(date '+%Y-%m-%d')

git add .
git reset -q -- '.env' '*.log' 'logs' '.vscode' 2>/dev/null
git commit -m "chore: update from ${date}"
git push
