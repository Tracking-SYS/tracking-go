#!/bin/bash

set -ex

PKGS=(
  github.com/grpc-ecosystem/grpc-gateway/v2
  google.golang.org/grpc/cmd/protoc-gen-go-grpc
  github.com/envoyproxy/protoc-gen-validate
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
  github.com/mitchellh/protoc-gen-go-json
  google.golang.org/protobuf/cmd/protoc-gen-go
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
)
for pk in "${PKGS[@]}"; do
  go get -u ${pk}
done