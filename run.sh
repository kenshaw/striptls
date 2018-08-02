#!/bin/bash

REMOTE=$1

if [ -z "$REMOTE" ]; then
  echo "usage: $0 <REMOTE>"
  exit 1
fi

docker run \
  --detach \
  --rm \
  --publish 8080:8080 \
  --name striptls \
  --env REMOTE="$REMOTE" \
  $USER/striptls:latest
