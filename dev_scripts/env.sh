#!/bin/bash

export POSTGRES_CONTAINER_NAME="postgres-dev"
export POSTGRES_USER="postgres"
export POSTGRES_PASSWORD="psql-dev-password"
export POSTGRES_DB="r9kd"
export POSTGRES_PORT="5432"
export POSTGRES_URL="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_CONTAINER_NAME:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable"

export R9KD_IMAGE_NAME="r9kd"
export R9KD_CONTAINER_NAME="r9kd"
export R9KD_PORT="8088"
