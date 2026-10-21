#!/usr/bin/env bash
set -euo pipefail

# Roda a partir da raiz do backend para que o .env seja encontrado
cd "$(dirname "$0")/.."

mkdir -p logs

# Mostra no terminal e grava no arquivo de log (stdout + stderr),
# sobrescrevendo o log da execução anterior
go run ./cmd/api "$@" 2>&1 | tee logs/api.log
