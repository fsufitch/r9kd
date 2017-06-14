#!/bin/bash
cd "${0%/*}"
source ./env.sh

echo Building image $R9KD_IMAGE_NAME...
docker build -t $R9KD_IMAGE_NAME \
  --build-arg R9KD_PORT=$R9KD_PORT \
  --build-arg POSTGRES_URL=$POSTGRES_URL \
  --build-arg ADMIN_KEY=$ADMIN_KEY \
  ..

echo Running $R9KD_IMAGE_NAME...
echo DB: $POSTGRES_URL
echo PORT: $R9KD_PORT
docker run -it --rm \
  --name $R9KD_CONTAINER_NAME \
  -p $R9KD_PORT:$R9KD_PORT \
  --link $POSTGRES_CONTAINER_NAME \
  $R9KD_IMAGE_NAME
