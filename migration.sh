#!/bin/bash
source .env

export MIGRATION_DSN="host=pg port=5432 dbname=$PG_DB_NAME user=$PG_USER password=$PG_PASSWORD"

# shellcheck disable=SC2153
sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v