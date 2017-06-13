#!/bin/bash
cd "${0%/*}"
source ./env.sh

docker run -d --name $POSTGRES_CONTAINER_NAME \
  -e POSTGRES_USER=$POSTGRES_USER \
  -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
  -e POSTGRES_DB=$POSTGRES_DB \
  postgres
