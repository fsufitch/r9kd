#!/bin/bash
cd "${0%/*}"
source ./env.sh

echo Building image $R9KD_IMAGE_NAME...
docker build -t $R9KD_IMAGE_NAME \
  --build-arg R9KD_PORT=$R9KD_PORT \
  --build-arg POSTGRES_URL=$POSTGRES_URL \
  --build-arg ADMIN_KEY=$ADMIN_KEY \
  ..

echo Running tests on $R9KD_IMAGE_NAME...
docker run -it --rm \
  --name $R9KD_TEST_CONTAINER_NAME \
  $R9KD_IMAGE_NAME \
    go test -cover \
    github.com/fsufitch/r9kd \
    github.com/fsufitch/r9kd/auth \
    github.com/fsufitch/r9kd/db \
    github.com/fsufitch/r9kd/model \
    github.com/fsufitch/r9kd/server \
    github.com/fsufitch/r9kd/server/channels \
    github.com/fsufitch/r9kd/server/message \
    github.com/fsufitch/r9kd/server/sender \
    github.com/fsufitch/r9kd/server/util
