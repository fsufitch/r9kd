#!/bin/bash
cd "${0%/*}"
source ./env.sh

docker build -t $R9KD_IMAGE_NAME ..
docker run -it --rm \
  --name $R9KD_CONTAINER_NAME \
  -e DATABASE_URL=$POSTGRES_URL \
  --link $POSTGRES_CONTAINER_NAME \
  $R9KD_IMAGE_NAME
