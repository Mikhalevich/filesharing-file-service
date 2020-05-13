#!/usr/bin/env bash

docker build -t file_service_app .
docker run -it --rm -p 50051:50051 file_service_app
