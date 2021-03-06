#!/usr/bin/env bash
set -x
set -eo pipefail

DB_USER="${POSTGRES_USER:=charlesonyewuenyi}" 
DB_PASSWORD="${POSTGRES_PASSWORD:=password}" 
DB_NAME="${POSTGRES_DB:=rest_api}" 
DB_PORT="${POSTGRES_PORT:=5432}"

# Allow to skip Docker if a dockerized Postgres database is already running
if [[ -z "${SKIP_DOCKER}" ]]; then
        docker run \
                -e POSTGRES_USER=${DB_USER} \
                -e POSTGRES_PASSWORD=${DB_PASSWORD} \
                -e POSTGRES_DB=${DB_NAME} \
                -p "${DB_PORT}":5432 \
                -d postgres \
                postgres -N 1000
fi

export PGPASSWORD="${DB_PASSWORD}"
"

