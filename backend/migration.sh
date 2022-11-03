# /bin/bash

source .env

migrate -path "$POSTGRES_MOGRATION_DIR" -database "$POSTGRES_SOURCE" -verbose "$1" $2