#! /usr/bin/env bash

protoc -I proto/ proto/file_server.proto --go_out=plugins=grpc:proto