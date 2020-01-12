#!/bin/bash
set -e

GOOS=linux go build
docker build . -t rails-init:latest
docker run rails-init:latest
docker run -it rails-init:latest /bin/bash
