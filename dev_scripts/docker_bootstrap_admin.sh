#!/bin/bash
cd "${0%/*}"
source ./env.sh

key=$1

if [ -z $key ]
then
  echo "Please pass in a key to add as admin"
  exit 1
fi

echo "INSERT INTO api_keys (key, admin) VALUES ('$key', TRUE);" > /tmp/bootstrap.sql;
echo "
  echo *:*:$POSTGRES_DB:$POSTGRES_USER:$POSTGRES_PASSWORD > ~/.pgpass;
  chmod 0600 ~/.pgpass;
  psql -a -d $POSTGRES_DB -U $POSTGRES_USER < bootstrap.sql;
  rm ~/.pgpass bootstrap.sql bootstrap.sh;
" > /tmp/bootstrap.sh

docker cp /tmp/bootstrap.sql $POSTGRES_CONTAINER_NAME:bootstrap.sql;
docker cp /tmp/bootstrap.sh $POSTGRES_CONTAINER_NAME:bootstrap.sh;

docker exec -t $POSTGRES_CONTAINER_NAME bash bootstrap.sh
