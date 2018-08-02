#!/bin/bash

GOOS=linux CGO_ENABLED=0 \
  go build \
  -ldflags='-w -s' \
  -o striptls

docker build -t $USER/striptls:latest .
